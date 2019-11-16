package parser

import (
	"strings"

	"github.com/KevinSmile/SQL-Image/fsql-img/query"
	"github.com/KevinSmile/SQL-Image/fsql-img/tokenizer"
)

var allAttributes = []string{"mode", "size", "time", "hash", "name"}

func isValidAttribute(attribute string) error {
	for _, valid := range allAttributes {
		if attribute == valid {
			return nil
		}
	}
	return &ErrUnknownToken{attribute}
}

// parseAttrs parses the list of attributes passed to the SELECT clause.
func (p *parser) parseAttrs(attributes *[]string, modifiers *map[string][]query.Modifier) error {
	for {
		ident := p.expect(tokenizer.Identifier)
		if ident == nil {
			return p.currentError()
		}

		if ident.Raw == "*" || ident.Raw == "all" {
			*attributes = allAttributes
		} else {
			p.current = ident

			attrModifiers := make([]query.Modifier, 0)
			attribute, err := p.parseAttr(&attrModifiers)
			if err != nil {
				return err
			}
			*attributes = append(*attributes, attribute.Raw)
			(*modifiers)[attribute.Raw] = attrModifiers
		}

		if p.expect(tokenizer.Comma) == nil {
			break
		}
	}
	return nil
}

// parseAttr recursively parses an attribute's modifiers and returns the
// associated attribute.
func (p *parser) parseAttr(modifiers *[]query.Modifier) (*tokenizer.Token, error) {
	ident := p.expect(tokenizer.Identifier)
	if ident == nil {
		return nil, p.currentError()
	}

	// ident is a modifier name (e.g. `FORMAT`) iff the next token is an open
	// paren, otherwise an attribute (e.g. `name`).
	if token := p.expect(tokenizer.OpenParen); token == nil {
		if err := isValidAttribute(ident.Raw); err != nil {
			return nil, err
		}
		return ident, nil
	}

	// In the case of chained modifiers, we want to recurse and parse each
	// inner modifier first. parseAttribute returns the associated attribute that
	// we're looking for.
	attribute, err := p.parseAttr(modifiers)
	if err != nil {
		return nil, err
	}
	if attribute == nil {
		return nil, p.currentError()
	}
	if err := isValidAttribute(attribute.Raw); err != nil {
		return nil, err
	}

	modifier := query.Modifier{
		Name:      strings.ToUpper(ident.Raw),
		Arguments: make([]string, 0),
	}

	// Parse the modifier arguments.
	for {
		if token := p.expect(tokenizer.Identifier); token != nil {
			modifier.Arguments = append(modifier.Arguments, token.Raw)
			continue
		}

		if token := p.expect(tokenizer.Comma); token != nil {
			continue
		}

		if token := p.expect(tokenizer.CloseParen); token != nil {
			*modifiers = append(*modifiers, modifier)
			return attribute, nil
		}
	}
}
