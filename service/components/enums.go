package components

const InternalServerError string = "{\"code\": 500, \"message\": \"Internal Server Error\"}"
const BadRequestError string = "{\"code\": 400, \"message\": \"Bad Request\"}"
const UnauthorizedError string = "{\"code\": 401, \"message\": \"Unauthorized\"}"
const NotFoundError string = "{\"code\": 404, \"message\": \"Not Found\"}"
const ConflictError string = "{\"code\": 409, \"message\": \"Conflict\"}"
const ForbiddenError string = "{\"code\": 403, \"message\": \"Forbidden\"}"

const InternalServerErrorF string = "{\"code\": 500, \"message\": \"Internal Server Error: %w\"}"
const BadRequestErrorF string = "{\"code\": 400, \"message\": \"Bad Request: %w\"}"
const UnauthorizedErrorF string = "{\"code\": 401, \"message\": \"Unauthorized: %w\"}"
const NotFoundErrorF string = "{\"code\": 404, \"message\": \"Not Found: %w\"}"
const ConflictErrorF string = "{\"code\": 409, \"message\": \"Conflict: %w\"}"
const ForbiddenErrorF string = "{\"code\": 403, \"message\": \"Forbidden: %w\"}"
