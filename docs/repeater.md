# Repeater

The repeater field allows adding and removing groups of fields dynamically via HTMX.

## Usage

```golang
fieldRepeater := form.NewRepeater(form.RepeaterOptions{
    Name: "addresses",
    Fields: []form.FieldInterface{
        form.NewStringField("street", "Street"),
        form.NewStringField("city", "City"),
    },
    Values: []map[string]string{
        {"street": "123 Main St", "city": "Springfield"},
        {"street": "456 Oak Ave", "city": "Shelbyville"},
    },
    RepeaterAddUrl:      "/repeater/add",
    RepeaterMoveUpUrl:   "/repeater/move-up",
    RepeaterMoveDownUrl: "/repeater/move-down",
    RepeaterRemoveUrl:   "/repeater/remove",
})
```

Each field in the repeater item is given a unique name, prefixed with the
repeater name, followed by the field name, and a dynamic index. This allows
each field to be identified when posting the form.

## Posted Data

The form when posted will produce:

```
addresses[street][] = 123 Main St
addresses[city][]   = Springfield
addresses[street][] = 456 Oak Ave
addresses[city][]   = Shelbyville
```
