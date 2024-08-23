function set_index_size(size, index) {
    let btngroup = document.getElementById('btn-group-index')
    let collapsables = document.getElementsByClassName('width-selectable')

    for (let i = 0; i < btngroup.children.length; i++) {
        if (i == index && !btngroup.children[i].classList.contains('active')) {
            btngroup.children[i].classList.add('active')
        }
        else if (btngroup.children[i].classList.contains('active')) {
            btngroup.children[i].classList.remove('active')
        }
    }

    for (let i = 0; i < collapsables.length; i++) {
        collapsables[i].children[0].classList.value = `col-md-${size}`
        collapsables[i].children[1].classList.value = `col-md-${(size < 10) ? 10 - size : 10}`
    }
}

function set_max_all() {
    let collapsables = document.getElementsByClassName('width-selectable')

    for (let i = 0; i < collapsables.length; i++) {
        let parent_cl = collapsables[i].parentElement.classList
        if (!parent_cl.contains('show')) {
            parent_cl.add('show')
        }
    }

    set_index_size(12)
}

function set_min_all() {
    let collapsables = document.getElementsByClassName('width-selectable')

    for (let i = 0; i < collapsables.length; i++) {
        let parent_cl = collapsables[i].parentElement.classList
        if (parent_cl.contains('show')) {
            parent_cl.remove('show')
        }
    }

    set_index_size(2)
}
