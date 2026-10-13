# Literary Lions Forum 
 
A Go web forum for a book club: SQLite-backed, cookie/UUID sessions, categorized
posts, like/dislike, filtering, and Docker deployment.

## Requirements

- Go 1.27+
- [Docker](https://docs.docker.com/get-started/get-docker/) (for containerized deployment)
- GCC (required if you use Windows machine, install [MinGW](https://www.mingw-w64.org/) for C compiler support)

## How To Run
Clone the repository and run the following commands:

```zsh
git clone https://gitea.kood.tech/hoangphuocvu/forum
cd forum
```

Docker deployment:
```zsh
docker compose up --build
```

Then website will be available at `http://localhost:8080`.

The first build will take ~1-2 minutes to complete.

Useful iteration commands:
```zsh
docker compose up           # start the conainer
docker compose down         # stop the container
docker compose restart      # restart the container
docker compose logs -f      # view the logs
```

Clean up unused Docker objects when iterating:

```zsh
docker image prune -f        # dangling images
docker container prune -f    # stopped containers
docker system df             # see what's using space
```

Or you can run it without Docker:
```zsh
go build -o forum
./forum
```
### ERD

![ERD Diagram](/erd.png)

## Extras

### Password Hashing

The forum uses sha256 for password hashing. When a user registers, their password is hashed using the sha256 algorithm before being stored in the database. This ensures that even if the database is compromised, the actual passwords remain secure. Powered by the `crypto/sha256` package in Go.

### Session Management

When a user logs in, a unique session ID (uuid) is generated and stored in a cookie on the user's browser. This session ID is also stored in the database, allowing the server to identify the user on subsequent requests. The session management system ensures that users remain logged in across different pages of the forum. Powered by the `uuid` package in Go since Go 1.27.0.

## Bonus Features

### Bcrypt Password Hashing

Hashing passwords with bcrypt is more secure than sha256. Bcrypt automatically handles salting and is designed to be computationally expensive, making it resistant to brute-force attacks. This feature is implemented by using the `golang.org/x/crypto/bcrypt` package.

It solves the problem of [rainbow table](https://www.huntress.com/cybersecurity-101/topic/rainbow-table-defined), which are precomputed lookup tables for matching the stolen hash (because 2 indentical passwords would produce the same hash). Since bcrypt includes a salt, it makes rainbow table attacks ineffective.

You can run it with ` --bcrypt` flag to enable bcrypt password hashing.

```zsh
./forum --bcrypt
```

### 404 Page
The 404 page is a simple soft landing page that shows the user that the page they are looking for does not exist if they navigate to a non-existent URL. The user can click the button to go back to the home page.

### Reset Password
Users can reset their passwords if they forgot it. The reset password link will be sent to their registered email address. Only registered email can receive the reset password link, but the error message will be generic to prevent the disclosure of registered email addresses. The link is valid for 30 minutes. Powered by `net/smtp` package.

You can try to register an account with your email to see the reset password feature in action. If you don't receive the reset link, please check your spam folder.

### Rate Limiter

The rate limiter powered by `github.com/phuocvu911/ratelimiter` package is used to prevent abuse of the forum by limiting the number of requests a user can make in a given time frame. This helps to protect the server from being overwhelmed by too many requests and ensures fair usage for all users.

### Search Feature

The forum includes a keyword-based search that lets users quickly find posts by matching their search query against post content and related metadata. Search results are paginated, so users can browse through larger result sets without overwhelming the page.

You can trigger a search from the posts view using the `q` query parameter, for example: `GET /posts/search?q=fiction`. Also there is advanced search criterias that can be explored in the website.

### Upload Images

Authenticated users can upload a profile image from their profile page. The uploaded image is saved in the database and served back when visiting the profile, with a default fallback image used when no custom avatar has been uploaded yet.

This feature supports common image formats and keeps the user’s uploaded picture available on future visits without requiring a separate file server.

### User Profile Page

Each signed-in user has a dedicated profile page that shows their account information, recent posts, comments, and liked content. The page also includes profile statistics so users can see how active they are across the forum.

From the profile, users can update their username, email, and password, and upload or replace their profile image at any time.