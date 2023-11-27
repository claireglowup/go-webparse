package delivery

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"errors"
	"html/template"
	"net/http"
	"net/url"
	"testfastprint/src/repository"

	"github.com/gin-gonic/gin"
)

type responseDataAPI struct {
	Error   int                                `json:"error,omitempty"`
	Ket     string                             `json:"ket,omitempty"`
	Version string                             `json:"version,omitempty"`
	Data    []repository.InsertProductParamsTx `json:"data,omitempty"`
}

func (d *delivery) InsertProduct(c *gin.Context) {

	formData := url.Values{
		"username": {"tesprogrammer281123C03"},           //berubah sewaktu waktu
		"password": {"4fa37776381c62c7ef8f896d54efc8ca"}, // berubah sewaktu waktu
	}

	req, err := http.NewRequest("POST", d.linkAPI, bytes.NewBufferString(formData.Encode()))
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	defer resp.Body.Close()

	var responseData responseDataAPI
	err = json.NewDecoder(resp.Body).Decode(&responseData)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	if responseData.Error != 0 {
		c.JSON(http.StatusBadRequest, errorResponse(errors.New(responseData.Ket)))
		return
	}

	err = d.usecase.InsertProductTx(c, responseData.Data)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
	})

}

func (d *delivery) GetAllProduct(c *gin.Context) {

	data, err := d.usecase.GetAllProduct(c)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, errorResponse(err))
			return
		}

		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	tmpl := template.Must(template.ParseFiles(
		"././view/frame.html",
		"././view/produk.html",
	))

	err = tmpl.ExecuteTemplate(c.Writer, "produk", data)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

}

func (d *delivery) GetProduk(c *gin.Context) {

	id := c.Param("id")

	data, err := d.usecase.GetProduct(c, id)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	tmpl := template.Must(template.ParseFiles(
		"././view/frame.html",
		"././view/edit.html",
	))

	err = tmpl.ExecuteTemplate(c.Writer, "edit", data)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

}

func (d *delivery) UpdateProduct(c *gin.Context) {

	id := c.Param("id")
	namaProduk := c.Request.FormValue("namaproduk")
	kategori := c.Request.FormValue("kategori")
	status := c.Request.FormValue("status")
	harga := c.Request.FormValue("harga")

	produk, err := d.usecase.GetProdukForUpdate(c, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	arg := repository.UpdateProductParamsTx{
		IdProduk:   id,
		NamaProduk: namaProduk,
		Kategori:   kategori,
		Status:     status,
		IdStatus:   produk.IDStatus,
		IdKategori: produk.IDKategori,
		Harga:      harga,
	}

	err = d.usecase.UpdateProductTx(c, arg)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "berhasil edit data",
	})

}

func (d *delivery) DeleteProduct(c *gin.Context) {

	id := c.Param("id")

	err := d.usecase.DeleteProduk(c, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "sukses menghapus",
	})

}
