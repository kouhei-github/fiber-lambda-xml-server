package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kouhei-github/fiber-sample-framework/repository"
	"github.com/kouhei-github/fiber-sample-framework/utils/authorization"
	"github.com/kouhei-github/fiber-sample-framework/utils/password"
	"strconv"
)

type loginBody struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func LoginHandler(c *fiber.Ctx) error {
	var payload loginBody
	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusUnauthorized).SendString(err.Error())
	}

	user := &repository.User{}
	users, err := user.FindByEmail(payload.Email)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": err.Error()})

	}

	if len(users) == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "存在しないメールアドレスです"})
	}

	user = &users[0]

	// passwordが正しいか確認
	ok := password.VerifyPassword(payload.Password, user.Password)
	if !ok {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "パスワードが正しくないです"})

	}
	str := strconv.FormatUint(uint64(user.ID), 10)
	token, err := authorization.CreateJwtToken(str)
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	res := signUpResponse{Token: token, Id: user.ID, Email: user.Email}
	return c.Status(fiber.StatusCreated).JSON(res)
}

type signUpResponse struct {
	Token string `json:"token"`
	Id    uint   `json:"id"`
	Email string `json:"email"`
}

func SignUpHandler(c *fiber.Ctx) error {
	var login loginBody
	if err := c.BodyParser(&login); err != nil {
		return c.Status(fiber.StatusUnauthorized).SendString(err.Error())
	}

	// パスワードのハッシュ化
	hashPassword, err := password.HashPassword(login.Password)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": err.Error()})
	}

	user := &repository.User{
		Email:    login.Email,
		Password: hashPassword,
	}
	users, err := user.FindByEmail(login.Email)

	// ユーザーが存在するか確認
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": err.Error()})
	}

	// すでにEmailアドレスが存在したらレスポンスを返す
	if len(users) != 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Emailアドレスはすでに存在します"})
	}

	// ユーザーの保存
	if err := user.Save(); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "保存できませんでした"})
	}

	// ユーザーIDを文字列に変換
	str := strconv.FormatUint(uint64(user.ID), 10)
	token, err := authorization.CreateJwtToken(str)
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	// レスポンス
	res := signUpResponse{Token: token, Id: user.ID, Email: user.Email}
	return c.Status(fiber.StatusCreated).JSON(res)
}
