Vue.component("display-pie", {
    props: {
        index: {
            type: Number,
            required: true
        },
        data: {
            type: Array,
            required: true
        }
    },
    template: "<div class=\"full\">\n\n\t\t<canvas class=\"full\" :ref=\"'pie-canvas-'+index\">Canvas not enabled</canvas>\n\t\n\t</div>",
    data() {
        return {
            canvas: {
                context: null
            },
            ctx: null,
            size: 0,
            p: 0,
            cx: 0,
            cy: 0
        };
    },
    methods: {
        onResize() {
            this.resizeCanvas();
            this.draw();
        },
        drawPlaceholder(msg) {
            if (!this.ctx || this.size < 10) {
                return;
            }
            this.ctx.fillStyle = '#EEE';
            this.ctx.font = Math.max(14, Math.round(2 * this.p)) + "px sans-serif";
            this.ctx.textAlign = 'center';
            this.ctx.textBaseline = 'middle';
            const lines = msg.split(' — ');
            let y = this.cy - (lines.length - 1) * 8;
            for (let i = 0; i < lines.length; i++) {
                this.ctx.fillText(lines[i], this.cx, y + i * 16);
            }
        },
        resizeCanvas() {
            const elem = this.$refs['pie-canvas-' + this.index];
            if (!elem) {
                return;
            }
            let box = elem.closest(".display-box");
            if (!box) {
                box = elem.parentElement;
            }
            let w = box ? box.clientWidth : 0;
            let h = box ? box.clientHeight : 0;
            if (w < 40 || h < 40) {
                w = Math.min(400, Math.max(280, window.innerWidth * 0.38));
                h = w;
            }
            elem.width = w;
            elem.height = h;
            this.size = Math.min(w, h);
            this.p = this.size / 100;
            this.cx = w / 2;
            this.cy = h / 2;
        },
        getPoint(g) {
            if (g < 0.25) {
                return {
                    "x": -1 + g * 8,
                    "y": -1
                };
            } else if (g < 0.5) {
                return {
                    "x": 1,
                    "y": -3 + g * 8
                };
            } else if (g < 0.75) {
                return {
                    "x": 5 - g * 8,
                    "y": 1
                };
            } else {
                return {
                    "x": -1,
                    "y": 7 - g * 8
                };
            }
        },
        draw() {
            const h = this.p;
            const canv = this.$refs['pie-canvas-' + this.index];
            const cw = canv ? canv.width : this.size;
            const ch = canv ? canv.height : this.size;
            this.ctx.fillStyle = '#000';
            this.ctx.fillRect(0, 0, cw, ch);
            if (!this.data || !this.data.length) {
                this.drawPlaceholder('Нет данных — загрузите логи и выберите игру');
                return;
            }
            const total = this.data[0] && this.data[0].amount;
            if (!total || total <= 0) {
                this.drawPlaceholder('Нет событий для этого режима');
                return;
            }
            this.ctx.textAlign = 'center';
            this.ctx.textBaseline = 'middle';
            this.ctx.font = "" + Math.round(2.5 * h) + "px Arial";
            let i = 0;
            for (let j = 1; j < this.data.length; j++) {
                const k = this.data[j];
                const l = k.amount / this.data[0].amount;
                const m = this.getPoint(i);
                const o = this.getPoint(i + l);
                this.ctx.beginPath();
                this.ctx.moveTo(this.cx, this.cy);
                this.ctx.lineTo(this.cx + m.x * 50 * h, this.cy + m.y * 50 * h);
                for (let q = 0; q < 4; q++) {
                    if (Math.floor(i * 4) + q < Math.floor((i + l) * 4)) {
                        const r = this.getPoint(Math.floor(i * 4 + 1 + q) / 4);
                        this.ctx.lineTo(this.cx + r.x * 50 * h, this.cy + r.y * 50 * h);
                    } else {
                        break;
                    }
                }
                this.ctx.lineTo(this.cx + o.x * 50 * h, this.cy + o.y * 50 * h);
                this.ctx.closePath();
                this.ctx.fillStyle = this.getColor(j);
                this.ctx.fillStyle = this.ctx.fillStyle + "DD";
                this.ctx.fill();
                i += l;
            }
            i = 0;
            for (let s = 1; s < this.data.length; s++) {
                const elem = this.data[s];
                const value = elem.amount / this.data[0].amount;
                const w = this.getPoint(i + value / 2);
                const z = {
                    "x": this.cx + w.x * 35 * h,
                    "y": this.cy + w.y * 35 * h
                };
                this.ctx.fillStyle = "rgba(0,0,0,0.5)";
                this.ctx.fillRect(z.x - 8 * h, z.y - 5 * h, 16 * h, 10 * h);
                this.ctx.fillStyle = '#FFF';
                this.ctx.fillText(elem.name, z.x, z.y - 3 * h);
                this.ctx.fillText(this.formatNumber(elem.amount), z.x, z.y);
                this.ctx.fillText(Math.round(value * 100) + " %", z.x, z.y + 3 * h);
                i += value;
            }
            this.ctx.fillStyle = '#000';
            this.ctx.fillRect(this.cx - 10 * h, this.cy - 10 * h, 20 * h, 20 * h);
            this.ctx.fillStyle = '#FFF';
            this.ctx.fillText("Total " + this.data[0].name, this.cx, this.cy - 6 * h);
            this.ctx.fillText(this.formatNumber(this.data[0].amount), this.cx, this.cy - 2 * h);
            this.ctx.fillText("Average", this.cx, this.cy + 2 * h);
            this.ctx.fillText(this.formatNumber(this.data[0].average), this.cx, this.cy + 6 * h);
        }
    },
    watch: {
        data: {
            handler() {
                this.$nextTick(() => {
                    this.resizeCanvas();
                    this.draw();
                });
            },
            deep: true
        }
    },
    computed: {},
    provide() {
        return {
            provide: this.canvas
        };
    },
    created() { },
    mounted() {
        const elem = this.$refs['pie-canvas-' + this.index];
        this.canvas.context = elem.getContext('2d');
        this.ctx = this.canvas.context;
        this.resizeCanvas();
        this.draw();
        window.addEventListener('resize', this.onResize);
    },
    beforeDestroy() {
        window.removeEventListener('resize', this.onResize);
    },
});