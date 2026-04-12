function cadastrar() {
  const nome      = document.getElementById('inp-nome').value.trim();
  const email     = document.getElementById('inp-email').value.trim();
  const senha     = document.getElementById('inp-senha').value;
  const confirmar = document.getElementById('inp-confirmar').value;
  const profissao = document.getElementById('inp-profissao').value;
  const erro      = document.getElementById('cadastro-erro');
  const sucesso   = document.getElementById('cadastro-sucesso');

  erro.style.display    = 'none';
  sucesso.style.display = 'none';

  if (!nome) {
    mostrarErro(erro, 'Nome completo é obrigatório.');
    return;
  }

  if (!email) {
    mostrarErro(erro, 'Email é obrigatório.');
    return;
  }

  const emailRegex = /^[^\s@]+@[^\s@]+\.[^\s@]+$/;
  if (!emailRegex.test(email)) {
    mostrarErro(erro, 'Informe um email com formato válido.');
    return;
  }

  if (!senha) {
    mostrarErro(erro, 'Senha é obrigatória.');
    return;
  }

  if (senha.length < 6) {
    mostrarErro(erro, 'A senha deve ter pelo menos 6 caracteres.');
    return;
  }

  if (senha !== confirmar) {
    mostrarErro(erro, 'As senhas não coincidem.');
    return;
  }

  if (!profissao) {
    mostrarErro(erro, 'Selecione uma profissão.');
    return;
  }

  sucesso.style.display = 'block';
  setTimeout(() => {
    window.location.href = 'index.html';
  }, 2000);
}

function mostrarErro(el, msg) {
  el.textContent    = msg;
  el.style.display  = 'block';
}
