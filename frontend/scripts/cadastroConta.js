
async function cadastrarContas(titulo, valor, token) {
    const url = "http://localhost:8000/auth/account";

    const dados = {
        title: titulo,
        amount: parseFloat(valor)
    }

    console.log(dados.amount);

    try{
        const resposta = await fetch(url, {
            method: 'POST',
            headers: {
                "Content-Type": "application/json",
                "Authorization": `Bearer ${token}`
                },
            body : JSON.stringify(dados)
        });

        const resultado = await resposta.json();

        if(resultado.error){
            return {
                "error": resultado.error,
            }
        }

        if(resposta.status === 201){
            return {
                "sucesso":"Conta registrada com sucesso",
            }
        }

    }catch(error){
        console.log(error);
    }
}

const cookie = document.cookie.split('=')[1];
const log = document.getElementById('logInfo');
const btnCadastrar = document.getElementById('btn-cadastrar');

if(!log){
    console.log('Componente log não encontrado no DOM');
}

if(!cookie){
    window.location.replace("../views/login.html");
}

btnCadastrar.addEventListener('click', async () => {
    const titulo = document.getElementById('inputTitle');
    const valor = document.getElementById('inputAmount');

    if(titulo.value === "" || valor.value === ""){
        log.textContent = "Título e valor são obrigatórios";
        return
    }

    const valorCerto = valor.value.replace(",", ".");

    resposta = await cadastrarContas(titulo.value, valorCerto, cookie);

    if(resposta.error){
        console.log(resposta.error);
        return
    }

    if(resposta.sucesso){
        log.textContent = "Conta cadastrada com sucesso";
        titulo.value = "";
        valor.value = "";
        return
    }
});

