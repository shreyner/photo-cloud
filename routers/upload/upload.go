package upload

import (
	"gopkg.in/macaron.v1"
	"photo-cloud/models"
	"github.com/jinzhu/gorm"
)

func GetUpload(ctx *macaron.Context) {
	ctx.Error(403)
}

func PostUpload(ctx *macaron.Context, db *gorm.DB) {


	fileimg, _, error := ctx.Req.FormFile("img")
	defer fileimg.Close();

	if error != nil {
		ctx.Error(403, "Error upload file")
	} else {


		img, err := models.SaveImage(fileimg, db)

		if !err {
			ctx.Redirect("/" + img.GetFullName())
		} else {
			ctx.Error(500, "Error server")
		}

	}


}