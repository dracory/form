# Advanced

## Configuring the HTML Textarea

The HTML area uses the Trumbowyg WYSIWYG editor. To use it you need to add the following to the web page:

```html
<script src="//ajax.googleapis.com/ajax/libs/jquery/3.3.1/jquery.min.js"></script>
<script>window.jQuery || document.write('<script src="js/vendor/jquery-3.3.1.min.js"><\/script>')</script>
<script src="trumbowyg/dist/trumbowyg.min.js"></script>
<link rel="stylesheet" href="trumbowyg/dist/ui/trumbowyg.min.css">
```

If you need special options for the editor, you can use the `Options` field:

```golang
form.NewHtmlAreaField("content", "Content").
    WithHelp("The content of this blog post.").
    WithOptions(form.FieldOption{
        Key: "config",
        Value: `{
            btns: [
                ['viewHTML'],
                ['undo', 'redo'],
                ['formatting'],
                ['strong', 'em', 'del'],
                ['link'],
                ['unorderedList', 'orderedList'],
                ['removeformat'],
                ['fullscreen'],
            ],
            autogrow: true,
            removeformatPasted: true,
        }`,
    })
```

## Legacy API

The original `NewForm` / `NewField` constructors with options structs are still fully supported:

```golang
f := form.NewForm(form.FormOptions{
    ID: "FormCustomerUpdate",
})
f.SetFields([]form.Field{
    {
        Label: "First Name",
        Name:  "first_name",
        Type:  form.FORM_FIELD_TYPE_STRING,
        Value: "John",
    },
})
html := f.Build().ToHTML()
```
