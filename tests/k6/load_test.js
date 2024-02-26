import http from 'k6/http';
import { sleep, check } from 'k6';
import { htmlReport } from "https://raw.githubusercontent.com/benc-uk/k6-reporter/main/dist/bundle.js";

export const options = {
  stages: [
    { duration: '1m', target: 150 },
    { duration: '2m', target: 500 },
    { duration: '2m', target: 700 },
    { duration: '30s', target: 0 },
  ],
};

const getRandomId = () => {
  const random = Math.random();
  const scaledRandom = random * 5 + 1;
  return Math.floor(scaledRandom);
}

export default function() {
  const randomIndex = Math.floor(Math.random() * 2);

  const getRequest = {
    method: "GET",
    url: `http://localhost:9999/clientes/${getRandomId()}/extrato`
  }

  const postRequest = {
    method: "POST",
    url:`http://localhost:9999/clientes/${getRandomId()}/transacoes`,
    body: JSON.stringify({
      "valor": 1000,
      "tipo" : randomIndex === 0 ? "c" : "d",
      "descricao" : "descricao"
    }),
    params: {
      headers: { 'Content-Type': 'application/json' }
    }
  }

  const response = http.batch([postRequest, getRequest])

  response.forEach(res => {
    check(res, { 'status OK': (r) => r.status == 200 || r.status == 422 });
  })

  sleep(200);
}

export function handleSummary(data) {
  const fileName = `./summary/summary-${Date.now()}.html`
  return {
    [fileName]: htmlReport(data),
  };
}
