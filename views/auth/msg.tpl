<div class="card-body">
    <div class="alert alert-{{ .Msg.AlertType}}" role="alert">
        {{ .Msg.AlertContent}}
    </div>
    <a target="_blank" href="{{ .Msg.Link}}" class="btn btn-{{ .Msg.ButtonType}}">{{ .Msg.ButtonContent}}</a>
</div>