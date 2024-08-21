var canvas = document.getElementById('editor-canvas')
var img    = document.getElementById('editor-img')
var edits_wrapper = document.getElementById('edits-wrapper')
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
        ctx.rect(rect.x,rect.y,rect.width,rect.height);
        if      (rect.selected) ctx.strokeStyle = 'cyan'
        else                    ctx.strokeStyle = 'gray';
        ctx.lineWidth = (rect.selected || rect.hover) ? 4 : 2;
        ctx.stroke();
    }
    if (editmode && !mousedown) {
        ctx.beginPath();
        var width = last_rect.to.x - last_rect.from.x;
        var height = last_rect.to.y - last_rect.from.y;
        ctx.rect(last_rect.from.x,last_rect.from.y,width,height);
        ctx.strokeStyle = 'black';
        ctx.lineWidth = 3;
        ctx.stroke();
    }
}

function toggle_edit_mode() {
    var edit_toggle = document.getElementById('edit-toggle')
    var edit_save = document.getElementById('edit-save')
    var edit_next = document.getElementById('edit-next')

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

function select_saved_rect(date_utc) {
    for (let i = 0; i < saved_rects.length; i++) {
        if (date_utc == saved_rects[i].date) {
            return saved_rects[i]
        }
    }
    return null
}

function mouseenter(id) {
    date_utc = id.split('-')[1]
    rect = select_saved_rect(date_utc)
    if (rect.selected) return

    rect_holder = document.getElementById(id)
    rect_holder.classList.remove('bg-white')
    rect_holder.classList.add('bg-light')
    rect_holder.classList.add('border-2')

    rect.hover = true
    reset_canvas()
}

function mouseleave(id) {
    date_utc = id.split('-')[1]
    rect = select_saved_rect(date_utc)
    if (rect.selected) return

    rect_holder = document.getElementById(id)
    rect_holder.classList.add('bg-white')
    rect_holder.classList.remove('bg-light')
    rect_holder.classList.remove('border-2')

    rect.hover = false
    reset_canvas()
}

function rect_click(id) {
}

function save_rect() {
    if (0 == mousedown_cnt) {
        alert("No new selection!")
        return
    }
    var new_rect = {
        date: new Date().getTime(),
        x: last_rect.from.x,
        y: last_rect.from.y,
        width:  last_rect.to.x - last_rect.from.x,
        height: last_rect.to.y - last_rect.from.y,
        hover: false,
        selected: false
    }
    saved_rects.push(new_rect)
    mousedown_cnt = 0

    // TODO: Send back to server
    // TODO: Reload from server?

    edits_wrapper.innerHTML += `
<div class="mb-2 p-2 bg-white rounded box-shadow border border-gray" onmouseenter="mouseenter('rect-${new_rect.date}')" onmouseleave="mouseleave('rect-${new_rect.date}')" id="rect-${new_rect.date}" onclick="rect_click('rect-${new_rect.date}')">
    <div>Position: {${new_rect.x} ${new_rect.x} ${new_rect.width} ${new_rect.height}}</div>

    <div class="dropdown mt-2">
        <input value=""
               class="form-control dropdown-toggle"
               id="dropdown-original-${new_rect.date}"
               data-bs-toggle="dropdown"
               aria-expanded="false">
        <ul class="dropdown-menu dropdown-menu-dark p-0" aria-labelledby="dropdown-original-${new_rect.date}" id="dropdown-original-${new_rect.date}-dd" style="width: 100%;">
        </ul>
    </div>

    <div class="dropdown mt-2">
        <input value="" class="form-control dropdown-toggle" id="dropdown-translated-${new_rect.date}" data-bs-toggle="dropdown" aria-expanded="false">
        <ul class="dropdown-menu dropdown-menu-dark p-0" aria-labelledby="dropdown-translated-${new_rect.date}" id="dropdown-translated-${new_rect.date}-dd" style="width: 100%;">
        </ul>
    </div>
</div>
`

    document.getElementById(`rect-${new_rect.date}`).addEventListener("focusin", () => select_rect(`rect-${new_rect.date}`))
    document.getElementById(`rect-${new_rect.date}`).addEventListener("focusout", () => deselect_rect(`rect-${new_rect.date}`))
}

function edit_save_rect() {
    save_rect()
    toggle_edit_mode()
}

function edit_next_rect() {
    save_rect()
    reset_canvas()
}

function select_rect(id) {
    date_utc = id.split('-')[1]
    rect = select_saved_rect(date_utc)
    if (rect.selected) return
    rect.selected = true

    rect_holder = document.getElementById(id)
    rect_holder.classList.remove('bg-white')
    rect_holder.classList.remove('bg-light')
    rect_holder.classList.add('border-2')
    rect_holder.classList.add('border-info')

    reset_canvas()
}

function deselect_rect(id) {
    date_utc = id.split('-')[1]
    rect = select_saved_rect(date_utc)
    rect.selected = false

    rect_holder = document.getElementById(id)
    rect_holder.classList.add('bg-white')
    rect_holder.classList.remove('border-2')
    rect_holder.classList.remove('border-info')

    reset_canvas()
}

function select_datalist_object(id, new_text) {
    document.getElementById(id).value = new_text
    childs = document.getElementById(id+'-dd').children
    for (let i = 0; i < childs.length; i++) {
        if (childs[i].children[0].innerText == new_text && !childs[i].children[0].classList.contains('active')) {
            childs[i].children[0].classList.add('active')
        }
        else if (childs[i].children[0].classList.contains('active')) {
            childs[i].children[0].classList.remove('active')
        }
    }
}

function select_original(id, new_text) {
    select_datalist_object(id, new_text)
    // TODO: Send back to server
}

function select_translated(id, new_text) {
    select_datalist_object(id, new_text)
    // TODO: Send back to server
}
