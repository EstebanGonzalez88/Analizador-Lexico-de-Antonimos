document.getElementById("curpForm").addEventListener("submit", async function(e) {
  e.preventDefault();

  const data = {
    nombre: document.getElementById("nombre").value,
    paterno: document.getElementById("paterno").value,
    materno: document.getElementById("materno").value,
    fecha: document.getElementById("fecha").value,
    sexo: document.getElementById("sexo").value,
    estado: document.getElementById("estado").value
  };

  const response = await fetch("/curp", {
    method: "POST",
    headers: {"Content-Type": "application/json"},
    body: JSON.stringify(data)
  });

  const result = await response.json();
  document.getElementById("resultado").innerText = "CURP: " + result.curp;
});