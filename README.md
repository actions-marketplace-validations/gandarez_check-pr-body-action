# Check Pull Request Body Action

This action checks if a string must contains or not in the pull request body.

## Github Environment Variables

- `GITHUB_REPOSITORY`

## Inputs

### pr_number

**Optional** Pull Request number.

### contains

**Optional** String that must contains in pull request body. Defaults to empty string.

### not_contains

**Optional** String that must not contains in pull request body. Defaults to empty string.

## Env

### GITHUB_TOKEN

GitHub token.

## Example usage

```yaml
uses: gandarez/check-pr-body-action@v1
  with:
    pr_number: ${{ github.event.number }}
    contains: "something"
  env:
    GITHUB_TOKEN: ${{ secrets.GITHUB_TOKEN }}
```
