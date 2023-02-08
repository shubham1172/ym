# ym
YAML manipulator CLI

`ym` is a CLI tool to manipulate YAML files. It allows you to define a set of operations to apply to a YAML file and execute them.

## Install

```bash
go install github.com/shubham1172/ym@latest
```

## Example

If you have a `foo.yaml` file with the following content:

```yaml
people:
  - name: John
    age: 42
  - name: Jane
    age: 36
  - name: Judy
    age: 34
```

Create a file `ops.yaml` with the following content:

```yaml
- input: foo.yaml # The file to apply the operations to, relative to the current directory.
  output: generated-foo.yaml # If this is not set, the file will be updated in place.
  operations:
    - op: update
      path: .people[0].age # Follows a yq style path.
      value: 43
    - op: update
      path: .people[1].children
      value:
        - name: Jack
          age: 10
        - name: Jill
          age: 8
    - op: delete
      path: .people[2]
```

Then run `ym`:

```bash
# Run ym with the ops.yaml file.
ym -file ops.yaml
# Or, pipe the ops.yaml file to ym.
cat ops.yaml | ym
```

The `foo.yaml` file will be updated to:

```yaml
people:
  - name: John
    age: 43
  - name: Jane
    age: 36
    children:
      - name: Jack
        age: 10
      - name: Jill
        age: 8
```

## Credits

This project uses https://github.com/mikefarah/yq for YAML manipulation.