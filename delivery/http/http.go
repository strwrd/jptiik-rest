package http

import (
	"context"
	"crypto/md5"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/strwrd/jptiik-rest/model"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	"github.com/patrickmn/go-cache"
	goCache "github.com/patrickmn/go-cache"

	"github.com/strwrd/jptiik-rest/usecase"
)

// Delivery interface
type Delivery interface {
	Start() error
	Stop() error
}

// handler server
type handler struct {
	server  *echo.Echo
	usecase usecase.Usecase
	cache   *goCache.Cache
}

// NewHandler creating server object
func NewHandler(usecase usecase.Usecase) Delivery {
	// Creating echo framework server
	server := echo.New()

	// Creating cache object with default 5 minutes expiration and purge cache 10 minutes
	cache := goCache.New(5*time.Minute, 10*time.Minute)

	// Return server handler object
	return &handler{
		server,
		usecase,
		cache,
	}
}

// Start server listening port
func (h *handler) Start() error {
	h.server.HideBanner = true

	// Using server recover middleware (prevent app from crash when panic)
	h.server.Use(middleware.Recover())

	// Using server logger middleware
	h.server.Use(middleware.Logger())

	// Using custom cache middleware
	h.server.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			// Creating key hash based on MD5 Algorithm
			hash := md5.New()
			io.WriteString(hash, fmt.Sprintf("%v-%v", c.Request().Method, c.Request().URL))
			keyHash := fmt.Sprintf("%x", hash.Sum(nil))

			log.Printf("URL : %v-%v\n", c.Request().Method, c.Request().URL)
			log.Println("Hash : ", keyHash)

			// Checking key on cache
			result, found := h.cache.Get(keyHash)
			if found {
				log.Println("Cache : Found")

				// if cache found then send result from cache
				c.JSON(http.StatusOK, result)
				return nil
			}

			log.Println("Cache : Not Found")
			// if cache doesn't exist, process actual request to create response
			if err := next(c); err != nil {
				c.Error(err)
			}

			// save response to cache
			h.cache.Add(keyHash, c.Get("result"), cache.DefaultExpiration)
			return nil
		}
	})

	// GET /archieves
	h.server.GET("/archieves", func(c echo.Context) error {
		// Create timeout request
		ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
		defer cancel()

		// Do usecase
		res, err := h.usecase.GetAllArchieve(ctx)
		if err != nil {
			e := echo.NewHTTPError(http.StatusInternalServerError, err.Error())
			return c.JSON(e.Code, e)
		}

		// Copy result to cache & Return request
		c.Set("result", res)
		return c.JSON(http.StatusOK, res)
	})

	// GET /journals  || GET /journals?archieveId=:val || GET /journals?title=:val || GET /journals?author=:val
	h.server.GET("/journals", func(c echo.Context) error {
		// Create timeout request
		ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
		defer cancel()

		// Checking query parameter
		if c.QueryParam("archieveId") != "" {
			// Do usecase
			res, err := h.usecase.GetJournalsByArchieveID(ctx, c.QueryParam("archieveId"))
			if err != nil {
				e := echo.NewHTTPError(http.StatusInternalServerError, err.Error())
				return c.JSON(e.Code, e)
			}

			// Copy result to cache & Return request
			c.Set("result", res)
			return c.JSON(http.StatusOK, res)
		} else if c.QueryParam("title") != "" {
			res, err := h.usecase.GetJournalsByTitle(ctx, c.QueryParam("title"))
			if err != nil {
				e := echo.NewHTTPError(http.StatusInternalServerError, err.Error())
				return c.JSON(e.Code, e)
			}
			// Copy result to cache & Return request
			c.Set("result", res)
			return c.JSON(http.StatusOK, res)
		} else if c.QueryParam("author") != "" {
			res, err := h.usecase.GetJournalsByAuthor(ctx, c.QueryParam("author"))
			if err != nil {
				e := echo.NewHTTPError(http.StatusInternalServerError, err.Error())
				return c.JSON(e.Code, e)
			}
			// Copy result to cache & Return request
			c.Set("result", res)
			return c.JSON(http.StatusOK, res)
		}

		// Do usecase
		res, err := h.usecase.GetAllJournal(ctx)
		if err != nil {
			e := echo.NewHTTPError(http.StatusInternalServerError, err.Error())
			return c.JSON(e.Code, e)
		}
		// Copy result to cache & Return request
		c.Set("result", res)
		return c.JSON(http.StatusOK, res)
	})

	// GET /archieve?archieveId=:val || GET /archieve?code=:val
	h.server.GET("/archieve", func(c echo.Context) error {
		// Create timeout request
		ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
		defer cancel()

		// Checking query parameter
		if c.QueryParam("archieveId") != "" {
			// Do usecase
			res, err := h.usecase.GetArchieveByArchieveID(ctx, c.QueryParam("archieveId"))
			if err != nil {
				if err == model.ErrDataNotFound {
					e := echo.NewHTTPError(http.StatusNotFound, err.Error())
					return c.JSON(e.Code, e)
				}
				e := echo.NewHTTPError(http.StatusInternalServerError, err.Error())
				return c.JSON(e.Code, e)
			}
			// Copy result to cache & Return request
			c.Set("result", res)
			return c.JSON(http.StatusOK, res)
		} else if c.QueryParam("code") != "" {
			// Do usecase
			res, err := h.usecase.GetArchieveByCode(ctx, c.QueryParam("code"))
			if err != nil {
				if err == model.ErrDataNotFound {
					e := echo.NewHTTPError(http.StatusNotFound, err.Error())
					return c.JSON(e.Code, e)
				}
				e := echo.NewHTTPError(http.StatusInternalServerError, err.Error())
				return c.JSON(e.Code, e)
			}
			// Copy result to cache & Return request
			c.Set("result", res)
			return c.JSON(http.StatusOK, res)
		} else {
			return c.JSON(http.StatusNotAcceptable, "unknown query parameter")
		}
	})

	// GET /journal/:id
	h.server.GET("/journal/:id", func(c echo.Context) error {
		// Create timeout request
		ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
		defer cancel()

		// Do usecase
		res, err := h.usecase.GetJournalByJournalID(ctx, c.Param("id"))
		if err != nil {
			if err == model.ErrDataNotFound {
				e := echo.NewHTTPError(http.StatusNotFound, err.Error())
				return c.JSON(e.Code, e)
			}
			e := echo.NewHTTPError(http.StatusInternalServerError, err.Error())
			return c.JSON(e.Code, e)
		}
		// Copy result to cache & Return request
		c.Set("result", res)
		return c.JSON(http.StatusOK, res)
	})

	return h.server.Start(":8080")
}

// Shutdown() ...
func (h *handler) Stop() error {
	// Create timeout process
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	return h.server.Shutdown(ctx)
}
