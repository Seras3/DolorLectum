/*
 * mockapp
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: v1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package httpapi

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// EchoMessageGet - 
func EchoMessageGet(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
