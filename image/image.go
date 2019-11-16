package image

const (
	RepositoriesJsonFile = "/var/lib/docker/image/overlay2/repositories.json"
	ImagedbSha256Dir = "/var/lib/docker/image/overlay2/imagedb/content/sha256"
	LayerdbSha256Dir = "/var/lib/docker/image/overlay2/layerdb/sha256"
	Overlay2Dir = "/var/lib/docker/overlay2"
)

type Image struct {
	tag string
	sha256 string
}

type Layer struct {

}