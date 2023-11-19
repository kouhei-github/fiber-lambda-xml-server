package controller

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/kouhei-github/fiber-sample-framework/repository"
	"strconv"
)

type userResponse struct {
	Id       uint   `json:"id"`
	Email    string `json:"email"`
	UserName string `json:"userName"`
}

func GetAllUserHandler(c *fiber.Ctx) error {
	// LocalsメソッドでMiddlewareで設定した値をコンテキストから取得する
	userId := c.Locals("myUserId") // これで取得できる
	fmt.Println("---------------------------")
	fmt.Println(userId)
	fmt.Println("---------------------------")

	// ユーザーIDをuintに変更可能かどうか
	valStr, ok := userId.(string)
	if !ok {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "You failed to convert to the string from interface"})
	}
	valUint, err := strconv.ParseUint(valStr, 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "You failed to convert to the uint from string"})

	}

	// ユーザーIDで検索
	user := repository.User{}
	if err := user.FindById(uint(valUint)); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": err.Error()})
	}
	res := userResponse{Email: user.Email, Id: user.ID, UserName: user.UserName}
	return c.Status(fiber.StatusOK).JSON(res)
}
