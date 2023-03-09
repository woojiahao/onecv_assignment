# onecv_assignment

Take home assignment for GovTech GDS OneClientView.

This project is deployed on DigitalOcean using Docker Compose. The link is as such: https://onecv.woojiahao.com.

Note that there is already sample data provided but you can add more teachers and students
using `POST https://onecv.woojiahao.com/api/students` and `POST https://onecv.woojiahao.com/api/teachers` where
the `POST` body is just an `email` field with a valid email.

## Deploying locally

Ensure that you have Docker installed.

Clone the repository

```bash
git clone https://github.com/woojiahao/onecv_assignment.git
cd onecv_assignment/
```

Edit the `.env.example` with your intended database configuration and rename the file to `.env` accordingly.

Use Docker Compose to deploy the backend locally

```bash
docker compose --env-file .env up
```

Note that if you renamed the `.env.example` to something other than `.env`, you need to change the argument
for `--env-file` as well. Additionally, you will need to edit `Dockerfile` to `mv <name> .env` on line 9.

Once Docker Compose builds the image and starts, the database will be automatically populated with sample data and you
can test the backend.

```bash
curl http://localhost:8080/api/commonstudents?teacher=teacherken@gmail.com
```