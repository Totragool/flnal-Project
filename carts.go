package carts


import (

   "net/http"

    "github.com/gin-gonic/gin"

   "library/config"

   "library/entity"

)


func GetAll(c *gin.Context) {


   var carts []entity.Cart


   db := config.DB()

   results := db.Find(&carts)

   if results.Error != nil {

       c.JSON(http.StatusNotFound, gin.H{"error": results.Error.Error()})

       return

   }

   c.JSON(http.StatusOK, carts)


}


func Get(c *gin.Context) {


   ID := c.Param("id")

   var cart entity.Cart


   db := config.DB()

   results := db.First(&cart, ID)

   if results.Error != nil {

       c.JSON(http.StatusNotFound, gin.H{"error": results.Error.Error()})

       return

   }

   if cart.ID == 0 {

       c.JSON(http.StatusNoContent, gin.H{})

       return

   }

   c.JSON(http.StatusOK, cart)


}

func Create(c *gin.Context) {
    var cart entity.Cart
    if err := c.ShouldBindJSON(&cart); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Bad request, unable to map payload"})
        return
    }
    db := config.DB()
    result := db.Create(&cart)
    if result.Error != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create cart"})
        return
    }
    c.JSON(http.StatusCreated, gin.H{"message": "Created successful", "cart": cart})
 }

func Update(c *gin.Context) {


   var cart entity.Cart


   CartID := c.Param("id")


   db := config.DB()

   result := db.First(&cart, CartID)

   if result.Error != nil {

       c.JSON(http.StatusNotFound, gin.H{"error": "id not found"})

       return

   }


   if err := c.ShouldBindJSON(&cart); err != nil {

       c.JSON(http.StatusBadRequest, gin.H{"error": "Bad request, unable to map payload"})

       return

   }


   result = db.Save(&cart)

   if result.Error != nil {

       c.JSON(http.StatusBadRequest, gin.H{"error": "Bad request"})

       return

   }


   c.JSON(http.StatusOK, gin.H{"message": "Updated successful"})

}


func Delete(c *gin.Context) {


   id := c.Param("id")

   db := config.DB()

   if tx := db.Exec("DELETE FROM carts WHERE id = ?", id); tx.RowsAffected == 0 {

       c.JSON(http.StatusBadRequest, gin.H{"error": "id not found"})

       return

   }

   c.JSON(http.StatusOK, gin.H{"message": "Deleted successful"})

}