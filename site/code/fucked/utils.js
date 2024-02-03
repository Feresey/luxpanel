Vue.mixin({
    'methods': {
        'getColor'(_0x5d6f50) {
            switch (_0x5d6f50) {
                case 0x0:
                    return '#EEE';
                case 0x1:
                    return '#7D0';
                case 0x2:
                    return '#FB1';
                case 0x3:
                    return '#0BC';
                case 0x4:
                    return '#F31';
                case 0x5:
                    return '#390';
                case 0x6:
                    return '#70D';
                case 0x7:
                    return '#E23';
                case 0x8:
                    return '#33E';
                case 0x9:
                    return '#721';
                case 0xa:
                    return '#046';
                case 0xb:
                    return '#FA5';
                case 0xc:
                    return '#E0E';
            }
        },
        'extract'(_0x57d96f, _0x7176cb, _0x4f31dc) {
            const _0x4abe6d = _0x57d96f.indexOf(_0x7176cb) + _0x7176cb.length;
            if (_0x4f31dc) {
                const _0x38ae34 = _0x57d96f.indexOf(_0x4f31dc, _0x4abe6d);
                return _0x57d96f.substring(_0x4abe6d, _0x38ae34);
            } else {
                return _0x57d96f.substring(_0x4abe6d);
            }
        },
        'extractTrim'(_0x115e19, _0x13ec23, _0x3e3ef0) {
            return this.extract(_0x115e19, _0x13ec23, _0x3e3ef0).trim();
        },
        'extractFrom'(_0x48fa0a, _0xeac3c1, _0x1157f2) {
            if (_0x1157f2) {
                const _0x335a3b = _0x48fa0a.indexOf(_0x1157f2, _0xeac3c1);
                return _0x48fa0a.substring(_0xeac3c1, _0x335a3b);
            } else {
                return _0x48fa0a.substring(_0xeac3c1);
            }
        },
        'extractNumber'(_0x39b67e, _0x37ec62, _0x9b6b28) {
            const _0x3e10f9 = _0x39b67e.indexOf(' ', _0x9b6b28);
            return parseInt(_0x39b67e.substring(_0x37ec62, _0x3e10f9).trim());
        },
        'formatNumber'(_0x195d54) {
            let _0x2f7f6e = '';
            if (_0x195d54 > 100) {
                _0x2f7f6e += Math.round(_0x195d54);
                let _0x42a94b = 3;
                while (_0x2f7f6e.length > _0x42a94b) {
                    const _0x5d6cd7 = _0x2f7f6e.length - _0x42a94b;
                    _0x2f7f6e = _0x2f7f6e.substring(0, _0x5d6cd7) + ' ' + _0x2f7f6e.substring(_0x5d6cd7);
                    _0x42a94b += 4;
                }
            } else {
                _0x2f7f6e += _0x195d54;
                _0x2f7f6e = _0x2f7f6e.substring(0, 4);
            }
            return _0x2f7f6e;
        },
        'formatTime'(_0x104f2c) {
            return Math.floor(_0x104f2c / 60) + ' min ' + Math.round(_0x104f2c % 0x3c) + ' sec';
        },
        'computeTime'(_0x5a2928) {
            const _0x2dfb7f = _0x5a2928.split(':');
            if (_0x2dfb7f.length === 3) {
                return parseInt(_0x2dfb7f[0]) * 3600 + parseInt(_0x2dfb7f[1]) * 60 + parseInt(_0x2dfb7f[2]);
            } else if (_0x2dfb7f.length === 2) {
                return parseInt(_0x2dfb7f[0]) * 60 + parseInt(_0x2dfb7f[1]);
            } else {
                return parseInt(_0x2dfb7f[0]);
            }
        },
        'buildPie'(_0x60bd5e, _0x407af6) {
            const _0x3bd269 = [{
                'name': _0x60bd5e,
                'amount': 0,
                'average': 0x0
            }];
            for (let _0x1c1ce9 = 0; _0x1c1ce9 < _0x407af6.length; _0x1c1ce9++) {
                _0x3bd269.push({
                    'name': _0x407af6[_0x1c1ce9],
                    'amount': 0x0
                });
            }
            return _0x3bd269;
        },
        'buildChart'(_0x5e3de1, _0x4ac91a) {
            const _0x39578a = [{
                'name': 'Average ' + _0x5e3de1,
                'data': [],
                'step': []
            }];
            for (let _0x2b3350 = 0; _0x2b3350 < _0x4ac91a.length; _0x2b3350++) {
                _0x39578a.push({
                    'name': _0x4ac91a[_0x2b3350],
                    'data': [],
                    'step': []
                });
            }
            return _0x39578a;
        },
        'buildTable'(_0x2d995e, _0x34ca59, _0x3ab681) {
            const _0x1a24e5 = [];
            _0x1a24e5.push(new Array(_0x3ab681.length + 2));
            for (let _0x340de6 = 0; _0x340de6 < _0x34ca59.length; _0x340de6++) {
                const _0x3d1443 = new Array(_0x3ab681.length + 2).fill(0);
                _0x3d1443[0] = _0x34ca59[_0x340de6];
                _0x1a24e5.push(_0x3d1443);
            }
            for (let _0xf349c4 = 0; _0xf349c4 < _0x3ab681.length; _0xf349c4++) {
                _0x1a24e5[0][_0xf349c4 + 1] = _0x3ab681[_0xf349c4];
            }
            _0x1a24e5.push(new Array(_0x3ab681.length + 2).fill(0));
            _0x1a24e5[0][0] = _0x2d995e;
            _0x1a24e5[_0x34ca59.length + 1][0] = 'Total';
            _0x1a24e5[0][_0x3ab681.length + 1] = 'Total';
            return _0x1a24e5;
        },
        'initProcess'(_0x449dc5, _0x48ff5c, _0x51bb35, _0x2a796c) {
            const _0x33d75c = _0x449dc5.game.teams;
            console.log(_0x48ff5c + ' computing have begun');
            _0x449dc5.pies = [];
            _0x449dc5.charts = [];
            _0x449dc5.tables = [];
            _0x449dc5.pies.push(this.buildPie(_0x48ff5c, _0x33d75c[0]));
            _0x449dc5.pies.push(this.buildPie(_0x48ff5c, _0x33d75c[1]));
            _0x449dc5.charts.push(this.buildChart(_0x48ff5c, _0x33d75c[0]));
            _0x449dc5.charts.push(this.buildChart(_0x48ff5c, _0x33d75c[1]));
            _0x449dc5.tables.push(this.buildTable(_0x48ff5c, _0x33d75c[0], _0x33d75c[_0x2a796c]));
            _0x449dc5.tables.push(this.buildTable(_0x48ff5c, _0x33d75c[1], _0x33d75c[_0x51bb35]));
        },
        'initCustom'(_0x1029a3, _0x5c2190, _0x2d1c77, _0x7b076c, _0x4b341c) {
            const _0x288b6a = _0x1029a3.game.teams;
            console.log(_0x5c2190 + ' computing have begun');
            _0x1029a3.pies = [];
            _0x1029a3.charts = [];
            _0x1029a3.tables = [];
            _0x1029a3.pies.push(this.buildPie(_0x5c2190, _0x2d1c77));
            _0x1029a3.pies.push(this.buildPie(_0x5c2190, _0x2d1c77));
            _0x1029a3.charts.push(this.buildChart(_0x5c2190, _0x2d1c77));
            _0x1029a3.charts.push(this.buildChart(_0x5c2190, _0x2d1c77));
            _0x1029a3.tables.push(this.buildTable(_0x5c2190, _0x288b6a[_0x7b076c], _0x2d1c77));
            _0x1029a3.tables.push(this.buildTable(_0x5c2190, _0x288b6a[_0x4b341c], _0x2d1c77));
        },
        'initLives'(_0x1a9703) {
            _0x1a9703.killIndex = 0;
            _0x1a9703.totalLives = [new Array(_0x1a9703.game.teams[0].length).fill(1), new Array(_0x1a9703.game.teams[1].length).fill(1)];
            _0x1a9703.teamLives = [_0x1a9703.game.teams[0].length, _0x1a9703.game.teams[1].length];
        },
        'fillProcess'(_0x4f0c69, _0x247560) {
            const _0x1dc5b9 = _0x4f0c69.game;
            const _0x3f5b7e = _0x1dc5b9.teams;
            _0x4f0c69.pies[_0x247560.side][0].amount += _0x247560.amount;
            _0x4f0c69.pies[_0x247560.side][0].average += _0x247560.amount / _0x3f5b7e[_0x247560.side].length;
            _0x4f0c69.pies[_0x247560.side][_0x247560.source + 1].amount += _0x247560.amount;
            const _0x521d5a = _0x4f0c69.charts[_0x247560.side][_0x247560.source + 1].step.length;
            const _0x45d630 = _0x4f0c69.charts[_0x247560.side][0].step.length;
            _0x4f0c69.charts[_0x247560.side][0].step[_0x45d630] = _0x247560.time - _0x1dc5b9.init;
            _0x4f0c69.charts[_0x247560.side][0].data[_0x45d630] = _0x4f0c69.pies[_0x247560.side][0].average;
            _0x4f0c69.charts[_0x247560.side][_0x247560.source + 1].step[_0x521d5a] = _0x4f0c69.charts[_0x247560.side][0].step[_0x45d630];
            _0x4f0c69.charts[_0x247560.side][_0x247560.source + 1].data[_0x521d5a] = _0x4f0c69.pies[_0x247560.side][_0x247560.source + 1].amount;
            _0x4f0c69.tables[_0x247560.side][_0x247560.source + 1][_0x247560.target + 1] += _0x247560.amount;
            _0x4f0c69.tables[_0x247560.side][_0x3f5b7e[_0x247560.side].length + 1][_0x247560.target + 1] += _0x247560.amount;
            _0x4f0c69.tables[_0x247560.side][_0x247560.source + 1][_0x3f5b7e[_0x247560.other].length + 1] += _0x247560.amount;
            _0x4f0c69.tables[_0x247560.side][_0x3f5b7e[_0x247560.side].length + 1][_0x3f5b7e[_0x247560.other].length + 1] += _0x247560.amount;
        },
        'fillCustom'(_0x47e965, _0x25fc2a, _0x2063bd) {
            const _0x51e1b0 = _0x47e965.game;
            const _0x160d41 = _0x51e1b0.teams;
            _0x47e965.pies[_0x25fc2a.side][0].amount += _0x25fc2a.amount;
            _0x47e965.pies[_0x25fc2a.side][0].average += _0x25fc2a.amount / _0x2063bd.length;
            _0x47e965.pies[_0x25fc2a.side][_0x25fc2a.type + 1].amount += _0x25fc2a.amount;
            const _0x52b5c7 = _0x47e965.charts[_0x25fc2a.side][_0x25fc2a.type + 1].step.length;
            const _0x4c8112 = _0x47e965.charts[_0x25fc2a.side][0].step.length;
            _0x47e965.charts[_0x25fc2a.side][0].step[_0x4c8112] = _0x25fc2a.time - _0x51e1b0.init;
            _0x47e965.charts[_0x25fc2a.side][0].data[_0x4c8112] = _0x47e965.pies[_0x25fc2a.side][0].average;
            _0x47e965.charts[_0x25fc2a.side][_0x25fc2a.type + 1].step[_0x52b5c7] = _0x47e965.charts[_0x25fc2a.side][0].step[_0x4c8112];
            _0x47e965.charts[_0x25fc2a.side][_0x25fc2a.type + 1].data[_0x52b5c7] = _0x47e965.pies[_0x25fc2a.side][_0x25fc2a.type + 1].amount;
            _0x47e965.tables[_0x25fc2a.side][_0x25fc2a.player + 1][_0x25fc2a.type + 1] += _0x25fc2a.amount;
            _0x47e965.tables[_0x25fc2a.side][_0x160d41[_0x25fc2a.side].length + 1][_0x25fc2a.type + 1] += _0x25fc2a.amount;
            _0x47e965.tables[_0x25fc2a.side][_0x25fc2a.player + 1][_0x2063bd.length + 1] += _0x25fc2a.amount;
            _0x47e965.tables[_0x25fc2a.side][_0x160d41[_0x25fc2a.side].length + 1][_0x2063bd.length + 1] += _0x25fc2a.amount;
        },
        'negatePartial'(_0x316587, _0x31f168) {
            const _0x361fa0 = _0x316587.game;
            const _0x53cfab = _0x361fa0.teams;
            _0x316587.totalLives[_0x31f168.other][_0x31f168.target]++;
            _0x316587.teamLives[_0x31f168.other]++;
            const _0x5a3b37 = (_0x316587.totalLives[_0x31f168.other][_0x31f168.target] - 1) / _0x316587.totalLives[_0x31f168.other][_0x31f168.target];
            const _0x38778c = (_0x316587.teamLives[_0x31f168.other] - 1) / _0x316587.teamLives[_0x31f168.other];
            _0x316587.pies[_0x31f168.other][0].average *= _0x38778c;
            _0x316587.pies[_0x31f168.other][_0x31f168.target + 1].amount *= _0x5a3b37;
            const _0x582461 = _0x316587.charts[_0x31f168.other][_0x31f168.target + 1].step.length;
            const _0x459b0a = _0x316587.charts[_0x31f168.other][0].step.length;
            _0x316587.charts[_0x31f168.other][0].step[_0x459b0a] = _0x31f168.time - _0x316587.game.init;
            _0x316587.charts[_0x31f168.other][0].data[_0x459b0a] = _0x316587.pies[_0x31f168.other][0].average;
            _0x316587.charts[_0x31f168.other][_0x31f168.target + 1].step[_0x582461] = _0x316587.charts[_0x31f168.other][0].step[_0x459b0a];
            _0x316587.charts[_0x31f168.other][_0x31f168.target + 1].data[_0x582461] = _0x316587.pies[_0x31f168.other][_0x31f168.target + 1].amount;
            for (let _0x2a50f2 = 0; _0x2a50f2 < _0x316587.game.teams[_0x31f168.side].length; _0x2a50f2++) {
                _0x316587.tables[_0x31f168.other][_0x31f168.target + 1][_0x2a50f2 + 1] *= _0x5a3b37;
            }
            _0x316587.tables[_0x31f168.other][_0x31f168.target + 1][_0x53cfab[_0x31f168.side].length + 1] = _0x316587.pies[_0x31f168.other][_0x31f168.target + 1].amount;
        },
        'fillPartial'(_0xfe3586, _0x3397fd) {
            const _0x40c398 = _0xfe3586.game;
            const _0x45ea7d = _0x40c398.teams;
            _0xfe3586.pies[_0x3397fd.side][0].amount += _0x3397fd.amount;
            _0xfe3586.pies[_0x3397fd.side][0].average += _0x3397fd.amount / _0xfe3586.teamLives[_0x3397fd.side];
            _0xfe3586.pies[_0x3397fd.side][_0x3397fd.source + 1].amount += _0x3397fd.amount / _0xfe3586.totalLives[_0x3397fd.side][_0x3397fd.source];
            const _0x165b2f = _0xfe3586.charts[_0x3397fd.side][_0x3397fd.source + 1].step.length;
            const _0xccbd2 = _0xfe3586.charts[_0x3397fd.side][0].step.length;
            _0xfe3586.charts[_0x3397fd.side][0].step[_0xccbd2] = _0x3397fd.time - _0xfe3586.game.init;
            _0xfe3586.charts[_0x3397fd.side][0].data[_0xccbd2] = _0xfe3586.pies[_0x3397fd.side][0].average;
            _0xfe3586.charts[_0x3397fd.side][_0x3397fd.source + 1].step[_0x165b2f] = _0xfe3586.charts[_0x3397fd.side][0].step[_0xccbd2];
            _0xfe3586.charts[_0x3397fd.side][_0x3397fd.source + 1].data[_0x165b2f] = _0xfe3586.pies[_0x3397fd.side][_0x3397fd.source + 1].amount;
            _0xfe3586.tables[_0x3397fd.side][_0x3397fd.source + 1][_0x3397fd.target + 1] += _0x3397fd.amount / _0xfe3586.totalLives[_0x3397fd.side][_0x3397fd.source];
            _0xfe3586.tables[_0x3397fd.side][_0x3397fd.source + 1][_0x45ea7d[_0x3397fd.other].length + 1] = _0xfe3586.pies[_0x3397fd.side][_0x3397fd.source + 1].amount;
        },
        'fillTotals'(_0x15ddfd) {
            const _0x5eb726 = _0x15ddfd.game;
            const _0x41b865 = [0, 0];
            for (let _0x58ff1f = 0; _0x58ff1f < 2; _0x58ff1f++) {
                for (let _0x493bcb = 0; _0x493bcb < _0x5eb726.teams[_0x58ff1f].length; _0x493bcb++) {
                    _0x41b865[_0x58ff1f] += _0x15ddfd.pies[_0x58ff1f][_0x493bcb + 1].amount;
                }
                _0x15ddfd.pies[_0x58ff1f][0].amount = _0x41b865[_0x58ff1f];
                _0x15ddfd.tables[_0x58ff1f][_0x5eb726.teams[_0x58ff1f].length + 1][_0x5eb726.teams[_0x58ff1f > 0 ? 0 : 1].length + 1] = _0x41b865[_0x58ff1f];
            }
        }
    }
});