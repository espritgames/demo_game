# Demo game sever

## API

### Get current user level

```
GET /user/<int:user_id>

{
    "id": <user_id>,
    "level": <level>
}
```

## Run

```
docker run -p 8080:8080 docker.pkg.github.com/espritgames/demo_game/demo_game:latest
```
