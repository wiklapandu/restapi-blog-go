# BLOG REST API
## Overview
ini adalah source code rest api untuk Blog Rest Api.
untuk mencobanya tinggal clone atau download

```cli
https://github.com/wiklapandu/restapi-blog-go.git
```

## Requirements
- Golang
- Gin

## Quickstart
sebelum menggunakan ada baiknya untuk mengubah port atau host dari `API` ini.

### Change Env

<table>
    <thead>
        <tr>
            <td>
                Key ENV
            </td>
            <td>
                Description
            </td>
            <td>
                Example
            </td>
        </tr>
    </thead>
    <tbody>
        <tr>
            <td>
                DB_HOST
            </td>
            <td>
                untuk mengubah host databases 
            </td>
            <td>
                DB_HOST=localhost
            </td>
        </tr>
        <tr>
            <td>
                DB_NAME
            </td>
            <td>
                untuk mengubah nama databases
            </td>
            <td>
                DB_NAME=dev_go_restapi
            </td>
        </tr>
        <tr>
            <td>
                DB_USER
            </td>
            <td>
                untuk mengubah username databases
            </td>
            <td>
                DB_USER=root
            </td>
        </tr>
        <tr>
            <td>
                DB_PASSWORD
            </td>
            <td>
                untuk mengubah password databases
            </td>
            <td>
                DB_PASSWORD=root
            </td>
        </tr>
        <tr>
            <td>
                SECRET_KEY
            </td>
            <td>
                secret key untuk JSON Web Token (JWT)
            </td>
            <td>
                SECRET_KEY="Hiram Bartoletti"
            </td>
        </tr>
    </tbody>
</table>


untuk menjalankan api ketik di cli
```cli
go run main.go
```

## How to Use
### Login User
- URL: `/login`
  - Method `POST`
  - Data Params
    - Required:
      - `email` type **string** `required|email`
      - `password` type **string** `required|email`
  - Success Response:
    - Code: `201` <br> Content: 
      ```json
        {
            "status": "success",
            "token":  "TOKEN_AUTH",
        }
      ```
  - Error Response:
    - Code: `400` <br> Content: 
      ```json
        {
          "errors": "Key: 'Password' Error:Field validation for 'Password' failed on the 'required' tag",
          "status": "errors"
        }
      ```
  - Fail Response:
    - Code: `400` <br> Content: 
      ```json
        {
			"status":  "fail",
			"message": "MESSAGE",
		}
      ```

### Login User
- URL: `/login`
  - Method `POST`
  - Data Params
    - Required:
      - `name` type **string** `required`
      - `email` type **string** `required|email`
      - `password` type **string** `required|match:repeat_password`
      - `repeat_password` type **string** `required`
  - Success Response:
    - Code: `201` <br> Content: 
      ```json
        {
            "status": "success",
            "token":  "TOKEN_AUTH",
        }
      ```
  - Error Response:
    - Code: `400` <br> Content: 
      ```json
        {
          "errors": "Key: 'Password' Error:Field validation for 'Password' failed on the 'required' tag",
          "status": "errors"
        }
      ```
  - Fail Response:
    - Code: `400` <br> Content: 
      ```json
        {
			"status":  "fail",
			"message": "MESSAGE",
		}
      ```

### Add Blog
- URL: `/blog`
  - Method `POST`
  - Data Params
    - `title` type **string** `required`
    - `desc` type **string** `required`
    - `cat` type **array** `required`
  - Header Required
    ```json
    {
        "Authorization": "TOKEN_AUTH"
    }
    ```
  - Success Response:
    - Code: `201` <br> Content: 
      ```json
        {
            "status": "success",
            "blog":  {
                "title": "BLOG_TITLE",
                "slug": "BLOG_SLUG",
                "desc": "BLOG_DESC",
                "user_id": "user_id",
                "categories": [],
            },
        }
      ```
  - Error Response:
    - Code: `400` <br> Content: 
      ```json
        {
          "errors": "Key: 'Title' Error:Field validation for 'Title' failed on the 'required' tag",
          "status": "errors"
        }
      ```
  - Fail Response:
    - Code: `400` <br> Content: 
      ```json
        {
			"status":  "fail",
			"message": "MESSAGE",
		}
      ```
- URL: `/blog/:id`
  - Method `PUT`
  - Data Params
    - `title` type **string** `required`
    - `desc` type **string** `required`
    - `cat` type **array** `required`
  - Header Required
    ```json
    {
        "Authorization": "TOKEN_AUTH"
    }
    ```
  - Success Response:
    - Code: `201` <br> Content: 
      ```json
        {
            "status": "success",
            "blog":  {
                "title": "BLOG_TITLE",
                "slug": "BLOG_SLUG",
                "desc": "BLOG_DESC",
                "user_id": "user_id",
                "categories": [],
            },
        }
      ```
  - Error Response:
    - Code: `400` <br> Content: 
      ```json
        {
          "errors": "Key: 'Title' Error:Field validation for 'Title' failed on the 'required' tag",
          "status": "errors"
        }
      ```
  - Fail Response:
    - Code: `400` <br> Content: 
      ```json
        {
			"status":  "fail",
			"message": "MESSAGE",
		}
      ```
- URL: `/blog/:id`
  - Method `DELETE`
- URL: `/blog/:slug`
  - Method `GET`
  - Success Response:
    - Code: `201` <br> Content: 
      ```json
        {
            "status": "success",
            "blog":  {
                  "title": "BLOG_TITLE",
                  "slug": "BLOG_SLUG",
                  "desc": "BLOG_DESC",
                  "user_id": "user_id",
                  "categories": [],
                }
        }
      ```
  - Error Response:
    - Code: `400` <br> Content: 
      ```json
        {
          "errors": "Key: 'Title' Error:Field validation for 'Title' failed on the 'required' tag",
          "status": "errors"
        }
      ```
  - Fail Response:
    - Code: `400` <br> Content: 
      ```json
        {
			"status":  "fail",
			"message": "MESSAGE",
		}
      ```
- URL: `/blog/`
  - Method `GET`
  - Method `GET`
  - Success Response:
    - Code: `201` <br> Content: 
      ```json
        {
            "status": "success",
            "blogs":  [
                {
                  "title": "BLOG_TITLE",
                  "slug": "BLOG_SLUG",
                  "desc": "BLOG_DESC",
                  "user_id": "user_id",
                  "categories": [],
                }
            ],
        }
      ```