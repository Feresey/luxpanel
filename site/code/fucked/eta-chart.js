Vue.component('eta-chart', {
    'props': {
        'index': {
            'type': Number,
            'required': true
        },
        'curves': {
            'type': Array,
            'required': true
        }
    },
    'template': `<canvas class="full" :ref="'chart-canvas-'+index">Canvas not enabled</canvas>`,
    'data'() {
        return {
            'canvas': {
                'context': null
            },
            'ctx': null,
            'width': 0,
            'height': 0,
            'mouse': {
                'in': true,
                'x': 500,
                'y': 0x64
            }
        };
    },
    'methods': {
        'mouseIn'(_0x231889) {
            this.mouse.in = true;
            this.mouse.x = _0x231889.offsetX;
            this.mouse.y = _0x231889.offsetY;
            this.drawChart();
        },
        'mouseMove'(_0x424e1b) {
            this.mouse.x = _0x424e1b.offsetX;
            this.mouse.y = _0x424e1b.offsetY;
            console.log(this.mouse.x, this.mouse.y);
            this.drawChart();
        },
        'mouseOut'() {
            this.mouse.in = true;
            this.drawChart();
        },
        'getX'(_0x33a02c, _0x55efce, _0x507724, _0x5757de) {
            return _0x5757de.L + _0x5757de.P + (_0x33a02c - _0x55efce) * _0x507724;
        },
        'getY'(_0x4bed09, _0x46f4d4, _0x2e65a2, _0x81fc2b, _0x42e6ca) {
            return _0x42e6ca - _0x81fc2b.B - _0x81fc2b.P - (_0x4bed09 - _0x46f4d4) * _0x2e65a2;
        },
        'getLastIndex'(_0xb2750, _0x2add3e) {
            let _0x4b4eb7 = -1;
            for (let _0x57960f = 0; _0x57960f < _0xb2750.step.length; _0x57960f++) {
                if (_0xb2750.step[_0x57960f] <= _0x2add3e) {
                    _0x4b4eb7 = _0x57960f;
                } else {
                    break;
                }
            }
            return _0x4b4eb7;
        },
        'drawChart'() {
            this.ctx.fillStyle = '#000';
            this.ctx.fillRect(0, 0, this.width, this.height);
            const _0xc86e64 = this.curves;
            if (!_0xc86e64 || !_0xc86e64.length) {
                return;
            }
            const _0x1c6e40 = this.width;
            const _0x31f4ca = this.height;
            if (_0x1c6e40 === 0 || _0x31f4ca === 0 || _0xc86e64.length === 0) {
                return;
            }
            const _0x3691e3 = {
                'T': 8,
                'L': 50,
                'B': 20,
                'R': 8,
                'P': 0x0
            };
            let _0xe581dd = -1;
            let _0x5c7f5e = -1;
            let _0xd5239c = null;
            let _0x2c423d = null;
            let _0x275778 = null;
            for (let _0x4b0658 = 0; _0x4b0658 < _0xc86e64.length; _0x4b0658++) {
                const _0x12f69a = _0xc86e64[_0x4b0658];
                if (_0x12f69a.step.length > 0) {
                    if (_0xd5239c === null) {
                        _0xd5239c = _0x12f69a.step[0] > 999999999;
                        _0x2c423d = _0x12f69a.step[0];
                        _0x275778 = _0x12f69a.step[_0x12f69a.step.length - 1];
                    }
                    const _0x5263d4 = _0x12f69a.step[0];
                    const _0x5e74e3 = _0x12f69a.step[_0xc86e64[0].step.length - 1];
                    if (_0x5263d4 <= _0x2c423d && _0x5263d4 !== null) {
                        _0x2c423d = _0x5263d4;
                    }
                    if (_0x5e74e3 > _0x275778 && _0x5e74e3 !== null) {
                        _0x275778 = _0x5e74e3;
                    }
                }
                for (let _0x4b6792 = 0; _0x4b6792 < _0x12f69a.data.length; _0x4b6792++) {
                    const _0x14c17b = _0x12f69a.data[_0x4b6792];
                    if (_0xe581dd === -1 || _0x14c17b < _0xe581dd && _0x14c17b !== null) {
                        _0xe581dd = _0x14c17b;
                    }
                    if (_0x5c7f5e === -1 || _0x14c17b > _0x5c7f5e && _0x14c17b !== null) {
                        _0x5c7f5e = _0x14c17b;
                    }
                }
            }
            this.ctx.lineWidth = this.width / 200;
            this.ctx.fillStyle = '#EEE';
            this.ctx.strokeStyle = '#EEE';
            this.ctx.beginPath();
            this.ctx.moveTo(_0x3691e3.L, _0x31f4ca - _0x3691e3.B + 4);
            this.ctx.lineTo(_0x3691e3.L, _0x3691e3.T);
            this.ctx.lineTo(_0x3691e3.L + 4, _0x3691e3.T + 4);
            this.ctx.lineTo(_0x3691e3.L - 4, _0x3691e3.T + 4);
            this.ctx.lineTo(_0x3691e3.L, _0x3691e3.T);
            this.ctx.moveTo(_0x3691e3.L - 4, _0x31f4ca - _0x3691e3.B);
            this.ctx.lineTo(_0x1c6e40 - _0x3691e3.R, _0x31f4ca - _0x3691e3.B);
            this.ctx.lineTo(_0x1c6e40 - _0x3691e3.R - 4, _0x31f4ca - _0x3691e3.B + 4);
            this.ctx.lineTo(_0x1c6e40 - _0x3691e3.R - 4, _0x31f4ca - _0x3691e3.B - 4);
            this.ctx.lineTo(_0x1c6e40 - _0x3691e3.R, _0x31f4ca - _0x3691e3.B);
            this.ctx.stroke();
            const _0x3b2340 = (_0x31f4ca - _0x3691e3.T - _0x3691e3.B - 2 * _0x3691e3.P) / (_0x5c7f5e - _0xe581dd);
            const _0x231e95 = (_0x1c6e40 - _0x3691e3.L - _0x3691e3.R - 2 * _0x3691e3.P) / (_0x275778 - _0x2c423d);
            for (let _0x19f2e5 = 0; _0x19f2e5 < _0xc86e64.length; _0x19f2e5++) {
                const _0x3a43b6 = _0xc86e64[_0x19f2e5];
                this.ctx.strokeStyle = _0x3a43b6.color;
                this.ctx.setLineDash([]);
                this.ctx.fillStyle = this.ctx.strokeStyle + '08';
                this.ctx.beginPath();
                this.ctx.lineTo(this.getX(_0x2c423d, _0x2c423d, _0x231e95, _0x3691e3), this.getY(_0xe581dd, _0xe581dd, _0x3b2340, _0x3691e3, _0x31f4ca));
                for (let _0x1e4186 = 0; _0x1e4186 < _0x3a43b6.data.length; _0x1e4186++) {
                    if (_0x3a43b6.data[_0x1e4186] !== null) {
                        this.ctx.lineTo(this.getX(_0x3a43b6.step[_0x1e4186], _0x2c423d, _0x231e95, _0x3691e3), this.getY(_0x3a43b6.data[_0x1e4186], _0xe581dd, _0x3b2340, _0x3691e3, _0x31f4ca));
                    }
                }
                this.ctx.lineTo(this.getX(_0x275778, _0x2c423d, _0x231e95, _0x3691e3), this.getY(_0x3a43b6.data[_0x3a43b6.data.length - 1], _0xe581dd, _0x3b2340, _0x3691e3, _0x31f4ca));
                this.ctx.stroke();
            }
            this.ctx.setLineDash([]);
            this.ctx.font = '12px Lato';
            this.ctx.fillStyle = '#EEE';
            this.ctx.textBaseline = 'top';
            this.ctx.textAlign = 'center';
            for (let _0x4992ec = Math.floor(_0x2c423d / 60) + 1; _0x4992ec < Math.floor(_0x275778 / 60) + 1; _0x4992ec++) {
                const _0x5c11c2 = _0x4992ec * 60;
                this.ctx.fillText(_0x4992ec + ' min', this.getX(_0x5c11c2, _0x2c423d, _0x231e95, _0x3691e3), _0x31f4ca - _0x3691e3.B + 8);
                this.ctx.fillText('|', this.getX(_0x5c11c2, _0x2c423d, _0x231e95, _0x3691e3), _0x31f4ca - _0x3691e3.B - 5);
            }
            this.ctx.textBaseline = 'middle';
            this.ctx.textAlign = 'right';
            const _0x24dddf = [5, 20, 100, 500, 2000, 10000, 50000, 200000, 1000000];
            let _0x2dbe2a = 1;
            for (let _0x23e159 = 0; _0x23e159 < _0x24dddf.length; _0x23e159++) {
                if (_0x5c7f5e - _0xe581dd > _0x24dddf[_0x23e159] * 4) {
                    _0x2dbe2a = _0x24dddf[_0x23e159];
                } else {
                    break;
                }
            }
            for (let _0x1b4dde = Math.floor(_0xe581dd / _0x2dbe2a) + 1; _0x1b4dde < Math.floor(_0x5c7f5e / _0x2dbe2a) + 1; _0x1b4dde++) {
                const _0x217096 = _0x1b4dde * _0x2dbe2a;
                this.ctx.fillText(this.formatNumber(_0x217096) + ' --', _0x3691e3.L + 4, this.getY(_0x217096, _0xe581dd, _0x3b2340, _0x3691e3, _0x31f4ca));
            }
            if (this.mouse.in) {
                let _0x396db7 = 0;
                let _0x38ddf3 = 0;
                let _0x49d133 = 0;
                let _0x3f45c4 = -1;
                for (let _0xac839d = 0; _0xac839d < _0xc86e64.length; _0xac839d++) {
                    const _0x2f5ac0 = _0xc86e64[_0xac839d];
                    for (let _0xd36ad7 = 0; _0xd36ad7 < _0x2f5ac0.step.length; _0xd36ad7++) {
                        const _0x46a4e1 = Math.abs(this.getX(_0x2f5ac0.step[_0xd36ad7], _0x2c423d, _0x231e95, _0x3691e3) - this.mouse.x);
                        if (_0x3f45c4 === -1 || _0x46a4e1 < _0x3f45c4) {
                            _0x3f45c4 = _0x46a4e1;
                            _0x396db7 = _0x2f5ac0.step[_0xd36ad7];
                            _0x38ddf3 = _0xd36ad7;
                        }
                    }
                }
                _0x3f45c4 = 999;
                this.ctx.strokeStyle = '#EEE';
                for (let _0x55ebc9 = 0; _0x55ebc9 < _0xc86e64.length; _0x55ebc9++) {
                    const _0x4cabdc = _0xc86e64[_0x55ebc9];
                    if (_0x4cabdc.step.length > 0) {
                        const _0x71352e = this.getLastIndex(_0x4cabdc, _0x396db7);
                        if (_0x71352e > -1) {
                            const _0x2b8d0d = Math.abs(this.getY(_0x4cabdc.data[_0x71352e], _0xe581dd, _0x3b2340, _0x3691e3, _0x31f4ca) - this.mouse.y);
                            if (_0x2b8d0d < _0x3f45c4) {
                                _0x3f45c4 = _0x2b8d0d;
                                _0x49d133 = _0x55ebc9;
                            }
                            if (_0x4cabdc.data[_0x71352e] !== null) {
                                this.ctx.beginPath();
                                this.ctx.arc(this.getX(_0x4cabdc.step[_0x71352e], _0x2c423d, _0x231e95, _0x3691e3), this.getY(_0x4cabdc.data[_0x71352e], _0xe581dd, _0x3b2340, _0x3691e3, _0x31f4ca), this.ctx.lineWidth * 1.5, 0, 2 * Math.PI);
                                this.ctx.stroke();
                            }
                        }
                    }
                }
                const _0x25039d = this.getLastIndex(_0xc86e64[_0x49d133], _0x396db7);
                if (_0x25039d > -1) {
                    const _0x14a199 = this.getX(_0xc86e64[_0x49d133].step[_0x25039d], _0x2c423d, _0x231e95, _0x3691e3);
                    const _0x27c39d = this.getY(_0xc86e64[_0x49d133].data[_0x25039d], _0xe581dd, _0x3b2340, _0x3691e3, _0x31f4ca);
                    this.ctx.beginPath();
                    this.ctx.arc(_0x14a199, _0x27c39d, this.ctx.lineWidth * 2.5, 0, 2 * Math.PI);
                    this.ctx.stroke();
                    this.ctx.lineWidth = 1;
                    this.ctx.beginPath();
                    this.ctx.moveTo(_0x3691e3.L, _0x27c39d);
                    this.ctx.lineTo(_0x1c6e40 - _0x3691e3.R, _0x27c39d);
                    this.ctx.moveTo(_0x14a199, _0x3691e3.T);
                    this.ctx.lineTo(_0x14a199, _0x31f4ca - _0x3691e3.B);
                    this.ctx.stroke();
                }
                const _0x5e34f8 = this.mouse.x < _0x1c6e40 / 2 ? _0x1c6e40 - 280 : _0x3691e3.L;
                this.ctx.font = '600 28px Arial';
                this.ctx.textBaseline = 'middle';
                this.ctx.textAlign = 'center';
                this.ctx.fillStyle = 'rgba(64, 60, 60, 0.25)';
                this.ctx.fillRect(_0x5e34f8, 0, 280, _0x31f4ca - _0x3691e3.B);
                this.ctx.fillStyle = 'rgba(255, 255, 255, 0.8)';
                const _0x4015bc = this.formatTime(_0xc86e64[_0x49d133].step[_0x25039d]);
                this.ctx.fillText(_0x4015bc, 150 + _0x5e34f8, 24);
                const _0x3c9a61 = [];
                let _0x31ef7a = '';
                for (let _0x2a9f73 = 0; _0x2a9f73 < _0xc86e64.length; _0x2a9f73++) {
                    const _0x5eebcb = this.getLastIndex(_0xc86e64[_0x2a9f73], _0x396db7);
                    if (_0x5eebcb > -1) {
                        if (_0x2a9f73 === _0x49d133) {
                            _0x31ef7a = _0xc86e64[_0x2a9f73].label;
                        }
                        _0x3c9a61.push({
                            'label': _0xc86e64[_0x2a9f73].name,
                            'color': _0xc86e64[_0x2a9f73].color,
                            'value': _0xc86e64[_0x2a9f73].data[_0x5eebcb]
                        });
                    }
                }
                _0x3c9a61.sort(function (_0xf0f7f7, _0xe0b3e4) {
                    return _0xf0f7f7.value < _0xe0b3e4.value ? 1 : _0xe0b3e4.value < _0xf0f7f7.value ? -1 : 0;
                });
                this.ctx.font = '600 18px Arial';
                for (let _0x57d42d = 0; _0x57d42d < _0x3c9a61.length; _0x57d42d++) {
                    this.ctx.textAlign = 'left';
                    this.ctx.fillStyle = _0x3c9a61[_0x57d42d].color;
                    this.ctx.fillStyle = _0x3c9a61[_0x57d42d].label === _0x31ef7a ? this.ctx.fillStyle + 'DD' : this.ctx.fillStyle + 'AA';
                    this.ctx.fillText(_0x3c9a61[_0x57d42d].label, _0x5e34f8 + 10, 50 + _0x57d42d * 20);
                    this.ctx.textAlign = 'right';
                    if (_0x3c9a61[_0x57d42d].value !== null) {
                        this.ctx.fillText(this.formatNumber(_0x3c9a61[_0x57d42d].value), _0x5e34f8 + 275, 50 + _0x57d42d * 20);
                    } else {
                        this.ctx.fillText('error', _0x5e34f8 + 275, 50 + _0x57d42d * 20);
                    }
                }
            }
        }
    },
    'watch': {
        'curves': function (_0x23a611, _0x3b388b) {
            this.drawChart();
        }
    },
    'computed': {},
    'provide'() {
        return {
            'provide': this.canvas
        };
    },
    'created'() { },
    'mounted'() {
        const _0x1fda97 = this.$refs['chart-canvas-' + this.index];
        this.canvas.context = _0x1fda97.getContext('2d');
        this.ctx = this.canvas.context;
        _0x1fda97.width = _0x1fda97.clientWidth;
        _0x1fda97.height = _0x1fda97.clientHeight;
        this.width = _0x1fda97.width;
        this.height = _0x1fda97.height;
        _0x1fda97.addEventListener('mouseover', _0x203d17 => {
            this.mouseIn(_0x203d17);
        });
        _0x1fda97.addEventListener('mousemove', _0x4fe45d => {
            this.mouseMove(_0x4fe45d);
        });
        _0x1fda97.addEventListener('mouseout', _0x4c26a8 => {
            this.mouseOut();
        });
        window.addEventListener('resize', _0x3fe892 => {
            _0x1fda97.width = _0x1fda97.clientWidth;
            _0x1fda97.height = _0x1fda97.clientHeight;
            this.width = _0x1fda97.width;
            this.height = _0x1fda97.height;
            this.drawChart();
        });
        this.drawChart();
    }
});