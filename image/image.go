package image

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"sort"
	"strings"
)

const (
	RepositoriesJsonFile = "/Users/Kevin/Downloads/dockerTest/var/lib/docker/image/overlay2/repositories.json"
	ImagedbSha256Dir = "/Users/Kevin/Downloads/dockerTest/var/lib/docker/image/overlay2/imagedb/content/sha256"
	LayerdbSha256Dir = "/Users/Kevin/Downloads/dockerTest/var/lib/docker/image/overlay2/layerdb/sha256"
	Overlay2Dir = "/Users/Kevin/Downloads/dockerTest/var/lib/docker/overlay2"
)

type Image struct {
	tag string
	sha256 string
}

type Layer struct {

}

type Repo struct {
	Repositories map[string]map[string]string
}

func GetImages() []Image{
	repositoriesJsonFile, _ := ioutil.ReadFile(RepositoriesJsonFile)
	var repo Repo
	var images []Image
	if err := json.Unmarshal(repositoriesJsonFile, &repo); err == nil {
		for _, imageTags := range repo.Repositories {
			for tag, sha256 := range imageTags {
				if !strings.Contains(tag, "@sha256:") {
					images = append(images, Image{tag, sha256[7:]})
				}
			}
		}
	} else {
		fmt.Println(err)
	}
	return images
}

func GetImageTags() []string{
	var tags []string
	for _, image := range GetImages() {
		tags = append(tags, image.tag)
	}
	return tags
}

func GetImageLayers(imageTag string) []string {
	var layers []string
	if dirs, err := ioutil.ReadDir(Overlay2Dir); err == nil {
		sort.Slice(dirs, func(i,j int) bool{
			return dirs[i].ModTime().After(dirs[j].ModTime())
		})
		for _, dir := range dirs {
			if dir.Name() != "l" {
				layers = append(layers, dir.Name())
			}
		}
	}
	return layers
}