const API = 'http://localhost:8000/v1';
const uid = localStorage.getItem('user_id');

function switchTab(type) {
    document.getElementById('form-login').classList.toggle('hidden', type === 'signup');
    document.getElementById('form-signup').classList.toggle('hidden', type === 'login');
    document.getElementById('tab-login').classList.toggle('active', type === 'login');
    document.getElementById('tab-signup').classList.toggle('active', type === 'signup');
}

async function login() {
    const res = await fetch(`${API}/login`, {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({
            username: document.getElementById('l-user').value,
            password: document.getElementById('l-pass').value
        })
    });
    const user = await res.json();
    if (user.id) {
        localStorage.setItem('user_id', user.id);
        location.href = 'notes.html';
    }
}

async function signup() {
    const res = await fetch(`${API}/signup`, {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({
            username: document.getElementById('s-user').value,
            password: document.getElementById('s-pass').value,
            email: document.getElementById('s-email').value,
            name: document.getElementById('s-name').value
        })
    });
    if (res.ok) switchTab('login');
}

function logout() {
    localStorage.clear();
    location.href = 'index.html';
}

async function loadNotes(pid = '') {
    const res = await fetch(`${API}/notes?user_id=${uid}&pid=${pid}`);
    const data = await res.json();
    return data.notes || [];
}

async function addNote(pid, input) {
    if (!input.value) return;
    const body = { user_id: uid, note: input.value };
    if (pid) body.parent_id = Number(pid);

    await fetch(`${API}/notes`, {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify(body)
    });
    renderNotes(pid ? document.querySelector(`[data-pid="${pid}"]`) : document.getElementById('tree'), pid);
}

async function toggleDone(id, currentDone, pid) {
    await fetch(`${API}/notes`, {
        method: 'PUT',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({ 
            user_id: uid, 
            note_id: Number(id), 
            done: currentDone == 1 ? 0 : 1 
        })
    });
    renderNotes(pid ? document.querySelector(`[data-pid="${pid}"]`) : document.getElementById('tree'), pid);
}

async function deleteNote(id, pid) {
    await fetch(`${API}/notes?user_id=${uid}&note_id=${Number(id)}`, { method: 'DELETE' });
    renderNotes(pid ? document.querySelector(`[data-pid="${pid}"]`) : document.getElementById('tree'), pid);
}

async function renderNotes(container, pid = '') {
    if (!container) return;
    container.innerHTML = '';
    container.setAttribute('data-pid', pid);
    const list = await loadNotes(pid);

    list.forEach(n => {
        const div = document.createElement('div');
        div.className = 'note-item';
        div.innerHTML = `
            <div class="note-row">
                <span class="note-text ${n.done == 1 ? 'done' : ''}">${n.note}</span>
                <button class="btn-s btn-done" onclick="toggleDone('${n.id}', ${n.done}, '${pid}')">V</button>
                <button class="btn-s" onclick="toggleChild(this, '${n.id}')">></button>
                <button class="btn-s btn-del" onclick="deleteNote('${n.id}', '${pid}')">×</button>
            </div>
            <div class="child-box hidden"></div>
        `;
        container.appendChild(div);
    });

    const addLine = document.createElement('div');
    addLine.className = 'add-line';
    addLine.innerText = '+ Добавить...';
    addLine.onclick = () => {
        const inp = document.createElement('input');
        inp.onkeydown = (e) => { if (e.key === 'Enter') addNote(pid, inp) };
        addLine.replaceWith(inp);
        inp.focus();
    };
    container.appendChild(addLine);
}

function toggleChild(btn, id) {
    const box = btn.parentElement.nextElementSibling;
    box.classList.toggle('hidden');
    if (!box.classList.contains('hidden')) renderNotes(box, id);
}

window.onload = async () => {
    const treeEl = document.getElementById('tree');
    const profileEl = document.getElementById('profile-data');

    if (treeEl && uid) renderNotes(treeEl);
    
    if (profileEl && uid) {
        const res = await fetch(`${API}/users/${uid}`);
        if (res.ok) {
            const u = await res.json();
            profileEl.innerHTML = `<p>Логин: ${u.username}</p><p>Имя: ${u.name}</p><p>Email: ${u.email}</p>`;
        } else {
            profileEl.innerText = "Ошибка загрузки профиля";
        }
    }
};