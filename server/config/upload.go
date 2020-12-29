package config

type Upload struct {
	ImgDomain			string `mapstructure:"img-domain" json:"imgUrl" yaml:"img-domain"`
	ImgUploadPath		string `mapstructure:"img-upload-path" json:"imgUploadPath" yaml:"img-upload-path"`
	ImgChunkPath		string `mapstructure:"img-chunk-path" json:"imgChunkPath" yaml:"img-chunk-path"`
	VideoDomain			string `mapstructure:"video-domain" json:"videoUrl" yaml:"video-domain"`
	VideoUploadPath		string `mapstructure:"video-upload-path" json:"videoUploadPath"  yaml:"video-upload-path"`
	VideoChunkPath		string `mapstructure:"video-chunk-path" json:"videoChunkPath" yaml:"video-chunk-path"`

}
