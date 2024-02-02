Vue.component('import', {
	props: {
	},
	template: `<div class="import row">

		<div class="import-box">
			<p>Pick your logs here</p>
			<input v-show="false" id="file" type="file" @change="LoadFiles($event)" multiple webkitdirectory/>
			<button class="button256" @click="searchFiles()">{{ready ? 'Logs loaded' : 'Pick Logs'}}</button>
		</div>

		<div v-show="!ready" class="import-box title-box">
			<p class="title">Log Analyzer</p>
			<p>for Star Conflict - by Eta</p>
		</div>
		
		<div v-show="!ready" class="import-box">
			<p>You don't have any logs ?</p>
			<a v-show="false" id="demo" href="./Demo" download="SC_logs_demo"></a>
			<button class="button256" @click="downloadLogs()">Download Logs</button>
		</div>

		<div v-show="ready" class="import-box">
			<p>Select your game</p>
			<select v-model="gameId" @change="gameChange($event)">
				<option v-for="(game, id) in games" :value="id">{{game.type  + ' (' + game.time + ')'}}</option>
			</select>
		</div>
		
		<div v-show="ready" class="import-box">
			<p>Chose a particular life</p>
			<select class="" v-model="lifeId" @change="lifeChange($event)" v-if="gameId">
				<option class="" v-for="(life, id) in games[gameId].lives" :value="id">{{life.ship  + ' (' + formatTime(life.duration) + ')'}}</option>
			</select>
		</div>
	
	</div>`,
	data() {
		return {
			ready: false,
			rawGame: [],
			rawCombat: [],
			selectedGame: [],
			selectedTeam: {},
			gameId: false,
			lifeId: false,
			games: [
				{
					id: 0,
					type: 'none',
					time: '14h20',
					duration: '2hour',
				},
				{
					id: 1,
					session: 41509120,
					type: 'example',
					map: 'nowhere',
					time: '14h20',
					duration: '10min',
					start: 99,
					end: 999,
					teams: [[], []],
					players: [
						{
							uuid: 38508,
							name: 'Eta',
							corp: 'luX',
							team: 2,
							main: true
						}
					],
					lives: [
						{
							id: 0,
							ship: 'all',
							time: '0min',
							duration: '10min',
							start: 99,
							end: 999,
						},
						{
							id: 1,
							ship: 'lightbringer',
							time: '0min',
							duration: '2min',
							start: 120,
							end: 250,
						}
					]
				}
			],
		}
	},
	methods: {
		searchFiles() {
			const loader = document.getElementById('file');
			this.ready = false;
			loader.click();
		},
		downloadLogs() {
			const download = document.getElementById('demo');
			download.click();
		},
		LoadFiles(event) {
			this.rawGame = [];
			this.rawCombat = [];
			for (let f = 0; f < event.target.files.length; f++) {
				const file = event.target.files[f];
				if (file.name === 'combat.log') {
					console.log('combat.log loaded');
					const reader = new FileReader();
					reader.onload = function (e) {
						this.rawCombat = e.target.result.split('\n');
						if (this.rawGame.length) {
							this.processData();
						}
					}.bind(this);
					reader.onerror = function (stuff) {
						console.log('combat error', stuff);
						//console.log(stuff.getMessage());
					};
					reader.readAsBinaryString(file);
				} else if (file.name === 'game.log') {
					console.log('game.log loaded');
					const reader = new FileReader();
					reader.onload = function (e) {
						this.rawGame = e.target.result.split('\n');
						if (this.rawCombat.length) {
							this.processData();
						}
					}.bind(this);
					reader.onerror = function (stuff) {
						console.log('game error', stuff);
						//console.log(stuff.getMessage());
					};
					reader.readAsBinaryString(file);
				}
			}
			this.gameId = 0;
		},
		processData() {
			this.ready = true;
			this.processGame();
			this.processCombat();
		},
		processGame() {
			const date = this.rawGame[2].substring(10, 20);
			const first = this.rawGame[4].substring(0, 5);
			const last = this.rawGame[this.rawGame.length - 2].substring(0, 5);
			this.games = [{
				id: 0,
				type: 'none selected',
				time: date + ' ' + first.split(':')[0] + 'h' + first.split(':')[1],
				duration: this.computeTime(last) - this.computeTime(first),
			}];
			let gameId = 0;
			let game = null;
			let state = 'game';
			for (let g = 0; g < this.rawGame.length; g++) {
				const line = this.rawGame[g];
				if (line.indexOf('MasterServerSession') > -1) {
					gameId++;
					this.games.push({
						id: gameId,
						session: this.extract(line, 'session ', ','),
						time: line.substring(0, 5),
						init: 0,
						teams: [[], []],
						players: [],
						lives: [{ ship: 'All lives' }],
					});
					game = this.games[gameId];
					state = 'players';
				}
				if (state === 'players') {
					if (line.indexOf('client: ADD_PLAYER') > -1) {
						if (this.extract(line, 'team ') !== '0') {
							const newPlayer = {
								uuid: this.extract(line, ', ', ')'),
								name: this.extract(line, '(', ' '),
								corp: this.extract(line, '[', ']'),
								team: parseInt(this.extract(line, 'team ')),
							};
							if (newPlayer.name.slice(-1) === ',') {
								newPlayer.name = newPlayer.name.slice(0, -1);
							}
							if (this.extract(line, 'status ', ' ') === '1') {
								newPlayer.main = true;
								game.mainTeam = newPlayer.team;
							}
							if (newPlayer.team !== 0) {
								game.players.push(newPlayer);
							}
						}
					} else if (line.indexOf('server assigned') > -1) {
						for (let p = 0; p < game.players.length; p++) {
							const player = game.players[p];
							if (player.main) {
								game.teams[0].unshift(player.name);
							} else if (player.team === game.mainTeam) {
								game.teams[0].push(player.name);
							} else {
								game.teams[1].push(player.name);
							}
						}
						state = 'game';
					}
				}
			}
		},
		processCombat() {
			let gameId = 0;
			let game = null;
			let lifeId = 0;
			let life = null;
			let state = 'lobby';
			const endLife = (line, index) => {
				const spawn = this.computeTime(game.lives[lifeId].time);
				const death = this.computeTime(line.substring(0, 8));
				game.lives[lifeId].duration = death - spawn;
				game.lives[lifeId].end = index;
			};
			for (let c = 0; c < this.rawCombat.length; c++) {
				const line = this.rawCombat[c];
				if (state === 'lobby') {
					if (line.indexOf('Connect to game') > -1) {
						state = 'game';
						gameId++;
						game = this.games[gameId];
						game.start = c;
						lifeId = 0;
						life = null;
						//let session = this.extract(line, 'session ', ' =');
						//let match = false;
						//while (!match) {
						//	gameId++;
						//	game = this.games[gameId];
						//	if (!game) {
						//		console.log('Fatal error, no game found for ID : '+session);
						//		match = true;
						//	}
						//	if (game.session === session) {
						//		match = true;
						//	}
						//}
					}
				}
				if (state === 'game') {
					if (line.indexOf('Start gameplay') > -1) {
						game.type = this.extract(line, 'gameplay \'', '\'');
						game.map = this.extract(line, 'map \'', '\'');
					}
					if (line.indexOf('Spawn SpaceShip') > -1) {
						if (!game.init) {
							game.init = this.computeTime(line.substring(0, 8));
						}
						if (this.extract(line, ' (', ',') === game.teams[0][0]) {
							lifeId++;
							game.lives.push({
								id: lifeId,
								ship: this.extract(line, '. \'', '\''),
								time: line.substring(0, 8),
								duration: null,
								start: c,
								end: 0,
							});
						}
					}
					if (line.indexOf('| Killed') > -1) {
						if (this.extract(line, 'Killed ', '\t') === game.teams[0][0]) {
							c++;
							while (this.rawCombat[c].indexOf('|    Participant') > -1) {
								c++;
							}
							endLife(line, c + 1);
						}
					}
					if (line.indexOf('Gameplay finished') > -1) {
						state = 'lobby';
						game.end = c;
						game.duration = parseInt(this.extract(line, 'game time ', '.'));
						game.lives[0].duration = game.duration;
						if (game.lives[lifeId].end === 0) {
							endLife(line, c);
						}
					}
				}
			}
		},
		gameChange(event) {
			this.lifeId = 0;
			if (this.gameId > 0) {
				const logs = this.rawCombat.slice(this.games[this.gameId].start, this.games[this.gameId].end);
				this.$root.$emit('import', this.games[this.gameId], logs);
				this.scrollDown();
			}
		},
		lifeChange(event) {
			let logs = [];
			if (this.lifeId > 0) {
				logs = this.rawCombat.slice(this.games[this.gameId].lives[this.lifeId].start, this.games[this.gameId].lives[this.lifeId].end);
			} else {
				logs = this.rawCombat.slice(this.games[this.gameId].start, this.games[this.gameId].end);
			}
			this.$root.$emit('import', this.games[this.gameId], logs);
			this.scrollDown();
		},
		scrollDown() {
			const scrollingElement = (document.scrollingElement || document.body);
			scrollingElement.scrollTop = scrollingElement.scrollHeight;
		}
	},
	computed: {

	},
	mounted() {

	}
});