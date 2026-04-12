function entrar() {
  const user = document.getElementById('inp-user').value.trim();
  const pass = document.getElementById('inp-pass').value;
  const erro = document.getElementById('login-erro');

  if (user === 'admin' && pass === '123456') {
    erro.style.display = 'none';
    sessionStorage.setItem('usuarioLogado', 'admin');
    sessionStorage.setItem('nomeUsuario', 'Analista N1');
    window.location.href = 'home.html';
  } else {
    erro.style.display = 'block';
  }
}
