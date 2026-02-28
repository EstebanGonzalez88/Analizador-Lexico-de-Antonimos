document.getElementById("curpForm").addEventListener("submit", async function(e) {

    e.preventDefault();

    const datos = {
        nombre: document.getElementById("nombre").value,
        paterno: document.getElementById("paterno").value,
        materno: document.getElementById("materno").value,
        fecha: document.getElementById("fecha").value,
        sexo: document.getElementById("sexo").value,
        estado: document.getElementById("estado").value
    };

    try {

        const response = await fetch("/curp", {
            method: "POST",
            headers: {
                "Content-Type": "application/json"
            },
            body: JSON.stringify(datos)
        });

        if (!response.ok) {
            const errorText = await response.text();
            document.getElementById("resultado").innerText = "Error: " + errorText;
            return;
        }

        const resultado = await response.json();
        document.getElementById("resultado").innerText = "CURP: " + resultado.curp;

    } catch (error) {
        document.getElementById("resultado").innerText = "Error de conexión con el servidor";
    }

});