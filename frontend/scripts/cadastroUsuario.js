
async function realizarCadastro(nome, email, senha){
    const url = "http://localhost:8000/user";
    const logInfo = document.getElementById('logInfo');

    if(!logInfo){
        console.log("Componente não encontrado no DOM");
        return
    }

    const dados = {
        name: nome, 
        email: email,
        password: senha
    }

    try{
        const resposta = await fetch(url, {
            method: "POST",
            headers: {
                'Content-Type':'application/json'
            },
            body: JSON.stringify(dados)
        });

         const resultado = await resposta.json();

        if(resposta.status === 201){
            logInfo.textContent = "Usuário cadastrado com sucesso. Você será redirecionado para a tela de Login...";

            setTimeout(() => {
                window.location.href = "../views/login.html";
            }, 3000);
        }

        if(resposta.status === 400){
            logInfo.textContent = resultado.error;
            return
        }

    }catch(error){
        console.log(`Error: ${error}`);
    }
}

const btnEnviar = document.getElementById('btn-enviar');

if(!btnEnviar){
    console.log("Elemento button não encontrado no DOM");
}

btnEnviar.addEventListener("click", () => {
    const nome = document.getElementById('inputNome').value;
    const email = document.getElementById('inputEmail').value;
    const senha = document.getElementById('inputSenha').value;

    const logInfo = document.getElementById('logInfo');

    if(!nome || !email || !senha) {
        console.log("Elementos não encontrado no DOM");
        return
    }

    if(nome === "" || email === "" || senha === ""){
        logInfo.textContent = "Todos os campos precisam estar preenchidos";
        return
    }

    realizarCadastro(nome, email, senha);
});