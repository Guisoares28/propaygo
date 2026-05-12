
async function buscarContas(){
    const url = "http://localhost:8000/auth/accounts";

    const token = document.cookie.split("=")[1];

    console.log(token);

    if(!token){
        window.location.replace("../views/login.html");
        return
    }

    const response = await fetch(url, {
        headers: {
            "Content-Type":"application/json",
            "Authorization":`Bearer ${token}`,
        },
    });
    const resultado = await response.json();

    if(resultado.error){
        console.log("Erro ao buscar as contas: " + resultado.error);
        return
    }

    return resultado.accounts

}

async function carregarContas() {

    const contas = await buscarContas();

   const tbody = document.getElementById("table_lines");

    contas.forEach(conta => {

    const tr = document.createElement("tr");

    tr.innerHTML = `
        <td>${conta.Title}</td>
        <td>R$ ${conta.Amount}</td>
        <td>${conta.Status}</td>
        <td>
            <button>Editar</button>
        </td>
    `;

    tbody.appendChild(tr);
});
}

carregarContas();