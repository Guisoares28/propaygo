

async function fazerLogin(email, senha) {
    const url = "http://localhost:8000/login";

    const dados = {
        email: email,
        password: senha
    };

    try {
        const resposta = await fetch(url, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify(dados)
        });

        if(resposta.ok){
            const resultado = await resposta.json();
            console.log("Login efetuado com sucesso: ",resultado)

            if(!resultado.token){
                console.log("Token não recebido");
                throw new Error("Token não enviado: ");
            }

            document.cookie = `token=${resultado.token}; path=/`;

            console.log("Cookie salvo vamos tentar redirecionar");

            window.location.replace('../views/home.html');

            console.log("A tentativa foi feita");

        }else{
            console.error("Erro no Login: ", resposta.statusText);
        }
    }catch(error){
        console.error("Erro de conexão: ",error);
    }
}

const btnEnviar = document.getElementById('btnEnviar');

if(!btnEnviar){
    throw new Error("botão não encontrado no DOM");
}

btnEnviar.addEventListener("click", () => {
    const email = document.getElementById('inputEmail').value;
    const senha = document.getElementById('inputSenha').value;
    const btn = document.getElementById('btnEnviar');
    btn.disabled = true;
    const logError = document.getElementById('logError');

    if(!btn){
        logError.textContent = "Botão não encontrado no DOM";
        return
    }

    if(email === "" || senha === ""){
        logError.textContent = "Email e Senha são obrigatórios";
        return        
    }

    fazerLogin(email, senha);

    btn.disabled = false;
});

