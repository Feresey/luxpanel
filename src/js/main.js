// Импортируйте наш пользовательский CSS
import '../scss/styles.scss'

// Импортируйте весь JS Bootstrap
import * as bootstrap from 'bootstrap'

import { CreateCharts } from './charts.js'

import './wasm_exec.js'

if (WebAssembly) {
    const go = new Go();
    WebAssembly.instantiateStreaming(fetch("gojs.wasm"), go.importObject).then((result) => {
        go.run(result.instance);
    });
} else {
    console.log("WebAssembly is not supported in your browser")
}

const pickLogs = document.getElementById('pick_logs');

pickLogs.addEventListener('change', function () {
    var data = {
        rawGame: "",
        rawCombat: "",
    }
    const files = this.files
    console.log("called", files);

    for (let f = 0; f < files.length; f++) {
        const file = files[f];

        if (file.name === 'combat.log') {
            const reader = new FileReader();
            reader.onload = function (e) {
                data.rawCombat = e.target.result;
                console.log('combat.log loaded', data.rawCombat.length);
                if (data.rawGame.length) {
                    parseFiles(data.rawGame, data.rawCombat);
                }
            }.bind(this);
            reader.onerror = function (stuff) { console.log('combat error', stuff); };
            reader.readAsBinaryString(file);
        } else if (file.name === 'game.log') {
            console.log('game.log loaded');
            const reader = new FileReader();
            reader.onload = function (e) {
                data.rawGame = e.target.result;
                console.log('combat.log loaded', data.rawGame.length);
                if (data.rawCombat.length) {
                    parseFiles(data.rawGame, data.rawCombat);
                }
            }.bind(this);
            reader.onerror = function (stuff) { console.log('game error', stuff); };
            reader.readAsBinaryString(file);
        }
    }
})

CreateCharts();