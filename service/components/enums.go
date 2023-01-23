package components

const InternalServerError string = "{\"code\": 500, \"message\": \"Internal Server Error\"}"
const BadRequestError string = "{\"code\": 400, \"message\": \"Bad Request\"}"
const UnauthorizedError string = "{\"code\": 401, \"message\": \"Unauthorized\"}"
const NotFoundError string = "{\"code\": 404, \"message\": \"Not Found\"}"
const ConflictError string = "{\"code\": 409, \"message\": \"Conflict\"}"
const ForbiddenError string = "{\"code\": 403, \"message\": \"Forbidden\"}"

const InternalServerErrorF string = "{\"code\": 500, \"message\": \"Internal Server Error: %s\"}"
const BadRequestErrorF string = "{\"code\": 400, \"message\": \"Bad Request: %s\"}"
const UnauthorizedErrorF string = "{\"code\": 401, \"message\": \"Unauthorized: %s\"}"
const NotFoundErrorF string = "{\"code\": 404, \"message\": \"Not Found: %s\"}"
const ConflictErrorF string = "{\"code\": 409, \"message\": \"Conflict: %s\"}"
const ForbiddenErrorF string = "{\"code\": 403, \"message\": \"Forbidden: %s\"}"
