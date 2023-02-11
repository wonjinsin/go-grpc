package controller

import (
	"fmt"
	"log"
	"net"
	"os"
	"phantom/util"
	"strconv"

	"github.com/labstack/echo/v4"
	"google.golang.org/grpc"
)

var zlog *util.Logger

type pickachuStatus struct {
	TRID       string      `json:"trid"`
	ResultCode string      `json:"resultCode"`
	ResultMsg  string      `json:"resultMsg"`
	ResultData interface{} `json:"resultData,omitempty"`
}

func init() {
	var err error
	zlog, err = util.NewLogger()
	if err != nil {
		log.Fatalf("InitLog module[controller] err[%s]", err.Error())
		os.Exit(1)
	}
}

// Init ...
func Init(port string) (server *grpc.Server, err error) {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", port))
	if err != nil {
		zlog.Errorw("failed to listen: %v", err)
		return nil, err
	}

	grpcServer := grpc.NewServer()
	log.Printf("start gRPC server on %s port", port)
	if err := grpcServer.Serve(lis); err != nil {
		zlog.Errorw("failed to serve: %s", err)
		return nil, err
	}

	return server, nil
}

func response(c echo.Context, code int, resultMsg string, result ...interface{}) error {
	strCode := strconv.Itoa(code)
	trid, ok := c.Request().Context().Value(util.TRID).(string)
	if !ok {
		trid = util.GetTRID()
	}

	res := pickachuStatus{
		TRID:       trid,
		ResultCode: strCode,
		ResultMsg:  resultMsg,
	}

	if result != nil {
		res.ResultData = result[0]
	}

	return c.JSON(code, res)
}

// UserController ...
type UserController interface {
	GetUser(c echo.Context) (err error)
	UpdateUser(c echo.Context) (err error)
	DeleteUser(c echo.Context) (err error)
}

// AuthController ...
type AuthController interface {
	Signup(c echo.Context) (err error)
	Signin(c echo.Context) (err error)
}
