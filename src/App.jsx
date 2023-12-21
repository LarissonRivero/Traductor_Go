import React, { useState } from 'react';
import Form from 'react-bootstrap/Form';
import 'bootstrap/dist/css/bootstrap.min.css';
import imagen from './assets/img/chatbot.jpg';
import './App.css';

const App = () => {
  const [inputText, setInputText] = useState('');
  const [translation, setTranslation] = useState('');
  const [targetLanguage, setTargetLanguage] = useState('es');

  const handleTranslate = async () => {
    try {
      const response = await fetch('http://localhost:8080/translate', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/x-www-form-urlencoded',
        },
        body: `text=${encodeURIComponent(inputText)}&targetLanguage=${encodeURIComponent(targetLanguage)}`,
      });
  
      const data = await response.json();
      setTranslation(data.translation);
    } catch (error) {
      console.error('Error al traducir:', error);
    }
  };

  return (
    <div className="card position-absolute top-50 start-50 translate-middle" style={{ width: '30rem', height: '39rem'}}>
      <img src={imagen} className="card-img-top" alt="Placeholder" />
      <div className="card-body">
        <h5 className="card-title">Traductor Creado con Go</h5>
        <p className="card-text">Facilita la traducción de texto del ingles al español, escribe la palabra o frase que quieras traducir y presiona Traducir</p>
        <label className='m-2'>
          Selecciona el idioma de destino:
          <select
            value={targetLanguage}
            onChange={(e) => setTargetLanguage(e.target.value)} className=''
          >
            <option value="es">Español</option>
            <option value="en">Inglés</option>
          </select>
        </label>
        <Form.Group className="mb-3" controlId="formBasicEmail">
        <Form.Control type="text"  className="text-center" placeholder="Ingrese el Texto o Frase" value={inputText}
        onChange={(e) => setInputText(e.target.value)}/>
      </Form.Group>
      <button className="btn btn-success" onClick={handleTranslate}>Traducir</button>
          <p className='m-4 fw-bold'>{translation}</p>
      </div>
    </div>
  );
};

export default App;