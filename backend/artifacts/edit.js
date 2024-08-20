var canvas = document.getElementById('editor-canvas')
var img    = document.getElementById('editor-img')
var edit_toggle = document.getElementById('edit-toggle')
var edit_save = document.getElementById('edit-save')
var edit_next = document.getElementById('edit-next')
var ctx    = canvas.getContext('2d')

var canvasx = canvas.offsetLeft
var canvasy = canvas.offsetTop
var mousex = 0
var mousey = 0
var mousedown = false
var mousedown_cnt = 0
var editmode = false
var last_rect = {
	from: { x: 0, y: 0},
	to:   { x: 0, y: 0}
}
var saved_rects = []
var scale = 1

canvas.width = img.width
canvas.height = img.height

ctx.drawImage(img, 0, 0)

function canvas_mousedown(e) {
    mousedown_cnt++
    scale = canvas.height / canvas.clientHeight
    canvasx = canvas.offsetLeft
    canvasy = canvas.offsetTop
    last_rect.from.x = (parseInt(e.clientX-canvasx) + window.scrollX) * scale
    last_rect.from.y = (parseInt(e.clientY-canvasy) + window.scrollY) * scale
    mousedown = true
}
function canvas_mouseup(e) {
    mousedown = false
    last_rect.to.x = (parseInt(e.clientX-canvasx) + window.scrollX) * scale
    last_rect.to.y = (parseInt(e.clientY-canvasy) + window.scrollY) * scale
}
function canvas_mousemove(e) {
    mousex = (parseInt(e.clientX-canvasx) + window.scrollX) * scale;
    mousey = (parseInt(e.clientY-canvasy) + window.scrollY) * scale;
    if(mousedown) {
        reset_canvas()
        ctx.beginPath();
        var width = mousex - last_rect.from.x;
        var height = mousey - last_rect.from.y;
        ctx.rect(last_rect.from.x,last_rect.from.y,width,height);
        ctx.strokeStyle = 'black';
        ctx.lineWidth = 3;
        ctx.stroke();
    }
}

function reset_canvas() {
    ctx.drawImage(img, 0, 0)
    for (let i = 0; i < saved_rects.length; i++) {
        let rect = saved_rects[i]
        ctx.beginPath();
        var width = rect.to.x - rect.from.x;
        var height = rect.to.y - rect.from.y;
        ctx.rect(rect.from.x,rect.from.y,width,height);
        ctx.strokeStyle = 'gray';
        ctx.lineWidth = 2;
        ctx.stroke();
    }
}

function toggle_edit_mode() {
    edit_toggle.blur()
    editmode = !editmode
    if (editmode) {
        canvas.onmousedown = canvas_mousedown
        canvas.onmouseup = canvas_mouseup
        canvas.onmousemove = canvas_mousemove
        canvas.style.borderLeft = "2px solid black"

        edit_toggle.innerText = "Cancel"
        edit_toggle.style.width = "32%"
        edit_toggle.classList.remove('btn-info')
        edit_toggle.classList.add('btn-warning')

        edit_save.hidden = false
        edit_next.hidden = false
    }
    else {
        canvas.onmousedown = null
        canvas.onmouseup = null
        canvas.onmousemove =  null
        canvas.style.border = ""

        edit_toggle.innerText = "Add"
        edit_toggle.style.width = "100%"
        edit_toggle.classList.remove('btn-warning')
        edit_toggle.classList.add('btn-info')

        edit_save.hidden = true
        edit_next.hidden = true
    }
    reset_canvas()
    mousedown_cnt = 0
}

function save_rect() {
    if (0 == mousedown_cnt) {
        alert("No new selection!")
        return
    }
    saved_rects.push({
        from: {
            x: last_rect.from.x,
            y: last_rect.from.y
        },
        to: {
            x: last_rect.to.x,
            y: last_rect.to.y
        }
    })
    mousedown_cnt = 0
}

function edit_save_rect() {
    save_rect()
    toggle_edit_mode()
}

function edit_next_rect() {
    save_rect()
    reset_canvas()
}
