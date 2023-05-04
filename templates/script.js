const criptomoedas = document.getElementById('criptomoedas');
const enviar = document.getElementById('enviar');

criptomoedas.addEventListener('change', () => {
  if (criptomoedas.value !== '') {
    enviar.removeAttribute('disabled');
  } else {
    enviar.setAttribute('disabled', true);
  }
});
