package main
type Image string
type Bucket string

type Storer interface {
	upload(image Image, bucket string)
	download(url string)
}

type AliStorer struct {

}

func (*AliStorer) upload(image Image, bucket string) {
	print("Ali upload\n")
}

func (*AliStorer) download(url string) {
	print("Ali Download\n")
}

type AwsStorer struct {

}

func (*AwsStorer) upload(image Image, bucket string) {
	print("Aws upload\n")
}

func (*AwsStorer) download(url string) {
	print("Aws Download\n")
}

func main() {
	var s Storer
	s  = new(AliStorer)
	s.upload("1.jpg", "Z01")
	s.download("s01.02")
	s  = new(AwsStorer)
	s.upload("1.jpg", "Z01")
	s.download("s01.02")
}

// 工厂模式后续可能要使用。。。