{{define "__PATH__"}}
  {{template "layout_form" .}}
{{end}}

{{define "content"}}
<form class="form-horizontal m" id="form-edit">
    <input type="hidden" name="id" value="{{.info.Id}}">
___REPLACE___
</form>
{{end}}

{{define "script"}}
<script>
    function submitHandler() {
        if ($.validate.form()) {
            $.operate.save(window.b5go.oesUrl, $('#form-edit').serialize());
        }
    }
</script>
{{end}}
