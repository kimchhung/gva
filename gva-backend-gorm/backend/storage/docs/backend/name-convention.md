# GoLand Naming Conventions

## Package Names

- Keep package names short, descriptive, and in lowercase, e.g., `articlefetcher`.
- Ensure the package name matches the directory name.
- Refrain from using acronyms in package names to maintain clarity.

## Struct, Function, Constants, Names

- Start struct names with a capital letter and use camel case, e.g., `ArticleFetcher`.
- Use nouns for struct names to clearly identify their purpose.
- Acronyms should be avoided to ensure readability.

## Variable Names

- Begin variable names with a lowercase letter and employ camel case, e.g., `articleFetcher`.
- Use nouns to give variables clear, descriptive names.
- Avoid using acronyms to maintain comprehensibility.

## Interface Names

- Interface names should start with a capital letter and use camel case, e.g., `ArticleFetcher`.
- Choose adjectives for interface names to describe behavior or capability.
- Avoid acronyms to ensure the name is self-explanatory.

## Tests

- Test files must end with a `_test.go` suffix for clarity and convention.
- Name test functions using the format `TestXxx`, where `Xxx` corresponds to the function being tested for consistency.

## API Endpoint

- Use a lowercase, dash-separated format for URLs, e.g., `[GET] /user/crawler-log`.
- Ensure endpoint names are descriptive and self-explanatory.

## Shorthand for Service and Repository

- Use fully descriptive names for services and repositories, e.g., `gameResultCacheService`.
- Common terms like service and repository may be abbreviated, e.g., `user_r` for `userRepository`, `user_s` for `userService`, but ensure clarity through context.
