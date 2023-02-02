# ym
YAML manipulator CLI

`ym` is a CLI tool to manipulate YAML files. It allows you to define a set of operations to apply to a YAML file and execute them.

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
- file: foo.yaml # The file to apply the operations to, relative to the current directory.
  operations:
    - op: replace
      path: .people[0].age # Follows a yq style path.
      value: 43
    - op: add
      path: .people[1].children
      value:
        - name: Jack
          age: 10
        - name: Jill
          age: 8
    - op: remove
      path: .people[2]
```

Then run `ym`:

```bash
ym -f ops.yaml
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