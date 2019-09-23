<template>
	<div>
		<div class="chat-enable">
			<img :src="logo ? logo : 'https://i.loli.net/2019/08/22/lOQSTxakG1NDMV9.png'" class="chat-logo" @click="enable = !enable" />
		</div>
		<div class="chat-wath" v-if="enable">
			<div class="wath-header">
				<div class="wath-header-f">
					<svg @click="menusEnable=!menusEnable" t="1566582106473" class="icon" viewBox="0 0 1024 1024" version="1.1" xmlns="http://www.w3.org/2000/svg" p-id="2227">
						<title>菜单</title>
						<path d="M832.2 264.6H192.3c-35.2 0-64-28.8-64-64s28.8-64 64-64h639.8c35.2 0 64 28.8 64 64 0.1 35.2-28.7 64-63.9 64zM832.2 577.9H196.1c-35.2 0-64-28.8-64-64s28.8-64 64-64h636.1c35.2 0 64 28.8 64 64s-28.8 64-64 64zM832.2 896.9H192.3c-35.2 0-64-28.8-64-64s28.8-64 64-64h639.8c35.2 0 64 28.8 64 64 0.1 35.2-28.7 64-63.9 64z" p-id="2228" />
					</svg>
					<div class="wath-header-menus" v-if="menusEnable">
						<a class="wath-theme" href="https://github.com/Carseason/c-chat" target="_blank">关于C聊</a>
					</div>
				</div>
				<div class="wath-header-t">
					<span>{{title}}</span>
				</div>
				<div class="wath-header-r">
					<svg t="1566374959239" class="icon" viewBox="0 0 1024 1024" version="1.1" xmlns="http://www.w3.org/2000/svg" p-id="5124" @click="enable = false">
						<title>关闭</title>
						<path d="M511.232 438.8352L112.9984 40.6016A51.2 51.2 0 0 0 40.6016 112.9984L438.784 511.232 40.6016 909.4656a51.2 51.2 0 1 0 72.3968 72.448l398.2336-398.2848 398.2336 398.2848a51.2 51.2 0 1 0 72.448-72.448l-398.2848-398.2336 398.2848-398.2336A51.2 51.2 0 0 0 909.4656 40.6016L511.232 438.784z" />
					</svg>
				</div>
			</div>
			<div class="wath-content" ref="showitems">
				<template v-for="value in ChatRecord">
					<li class="wath-record-list" :key="value.Date">
						<template v-if="value.Author.Name == Author.Name">
							<div class="wath-record-list-author">
								<div class="placeholder"></div>
								<div class="author-date">{{value.Date | dateForm}}</div>
								<div class="author-nickname" :title="value.Author.Name">{{value.Author.Name}}</div>
								<div class="author-avatar">
									<img :src="value.Author.Avatar" class="avatar" />
								</div>
							</div>
							<div class="wath-record-list-contens itself">
								<p v-html="value.Content"></p>
							</div>
						</template>
						<template v-else>
							<div class="wath-record-list-author">
								<div class="author-avatar">
									<img :src="value.Author.Avatar" class="avatar" />
								</div>
								<div class="author-nickname" :title="value.Author.Name">{{value.Author.Name}}</div>
								<div class="author-date">{{value.Date | dateForm}}</div>
							</div>
							<div class="wath-record-list-contens they">
								<p v-html="value.Content"></p>
							</div>
						</template>
					</li>
				</template>
			</div>
			<div class="wath-footer">
				<form method="post" @submit.prevent="UpdateContent()" class="wath-form">
					<div class="wath-edit">
						<div class="edit-textarea">
							<textarea placeholder="说点儿什么吧" v-model.trim="content" ref="editarea"></textarea>
						</div>
						<button type="submit" class="edit-push">
							<svg t="1566393222500" class="icon" viewBox="0 0 1025 1024" version="1.1" xmlns="http://www.w3.org/2000/svg" p-id="2215">
								<title>发送</title>
								<path
									d="M786.856169 1024a37.243789 37.243789 0 0 1-19.884397-5.681256L509.105874 852.615444l-137.612645 142.347025a36.296913 36.296913 0 0 1-39.453167 8.83751 36.928164 36.928164 0 0 1-23.6719-33.14066l-7.575008-259.1284-280.590922-141.40015A36.612539 36.612539 0 0 1 0.00221 536.990109a37.243789 37.243789 0 0 1 19.568771-31.562533L969.918862 4.53017a36.612539 36.612539 0 0 1 40.084418 3.156254 37.243789 37.243789 0 0 1 13.887515 36.928164l-201.053338 950.347881a35.981288 35.981288 0 0 1-21.778148 26.196903 36.928164 36.928164 0 0 1-14.20314 2.840628zM505.002745 767.396603a35.034412 35.034412 0 0 1 19.884396 5.996882l238.612753 154.340788 173.278308-820.625869L117.099209 536.990109l212.731476 106.365738L717.73422 260.502316a36.612539 36.612539 0 1 1 51.44693 52.078181L374.018231 702.377784l4.73438 177.697064 98.159479-101.631358a36.612539 36.612539 0 0 1 28.090655-11.046887z"
									p-id="2216"
								/>
							</svg>
						</button>
						<div class="edit-emoji">
							<div class="emoji-item" :class="{ on : emojiEnable}">
								<ul class="emoji-box">
									<a title="大笑" @click="UpdateCommentsContent('[大笑]')" class="emoji-list">
										<img src="https://i.loli.net/2019/08/05/wbmvoD5ycFJ79KA.gif" />
									</a>
									<a title="喷血" @click="UpdateCommentsContent('[喷血]')" class="emoji-list">
										<img src="https://i.loli.net/2019/08/05/EbY8FCLmKcQIWHz.gif" />
									</a>
									<a title="抠鼻" @click="UpdateCommentsContent('[抠鼻]')" class="emoji-list">
										<img src="https://i.loli.net/2019/08/05/xfhvsjE284nPVip.gif" />
									</a>
									<a title="吐" @click="UpdateCommentsContent('[吐]')" class="emoji-list">
										<img src="https://i.loli.net/2019/08/05/98nzCEk4pHriBV1.gif" />
									</a>
									<a title="微笑" @click="UpdateCommentsContent('[微笑]')" class="emoji-list">
										<img src="https://i.loli.net/2019/08/05/YoNItSivPL2ZJAn.gif" />
									</a>
									<a title="笑哭" @click="UpdateCommentsContent('[笑哭]')" class="emoji-list">
										<img src="https://i.loli.net/2019/08/05/PX6M1eFxrYVZNdC.gif" />
									</a>
									<a title="斜眼笑" @click="UpdateCommentsContent('[斜眼笑]')" class="emoji-list">
										<img src="https://i.loli.net/2019/08/05/vcDieIRrf25KlAz.gif" />
									</a>
									<a title="阴险" @click="UpdateCommentsContent('[阴险]')" class="emoji-list">
										<img src="https://i.loli.net/2019/08/05/tNqeoXlZcywU41R.gif" />
									</a>
									<a title="doge" @click="UpdateCommentsContent('[doge]')" class="emoji-list">
										<img src="https://i.loli.net/2019/08/05/sRAJywb56qphYK1.gif" />
									</a>
									<a title="狗头" @click="UpdateCommentsContent('[狗头]')" class="emoji-list">
										<img src="https://i.loli.net/2019/08/05/Ww6GxbhD1Vamq8A.gif" />
									</a>
									<a title="猪头" @click="UpdateCommentsContent('[猪头]')" class="emoji-list">
										<img src="https://i.loli.net/2019/08/05/jGbMDPfFphlrXuz.gif" />
									</a>
									<a title="马" @click="UpdateCommentsContent('[马]')" class="emoji-list">
										<img src="https://i.loli.net/2019/08/05/MnKwdZi4eBVj3LD.gif" />
									</a>
									<a title="牛" @click="UpdateCommentsContent('[牛]')" class="emoji-list">
										<img src="https://i.loli.net/2019/08/05/AguMSPy1DTN7BVt.gif" />
									</a>
									<a title="啤酒" @click="UpdateCommentsContent('[啤酒]')" class="emoji-list">
										<img src="https://i.loli.net/2019/08/05/q8i1MQUCE5YrOxF.gif" />
									</a>
								</ul>
							</div>
							<svg @click="emojiEnable = !emojiEnable" t="1566367724185" class="icon" viewBox="0 0 1024 1024" version="1.1" xmlns="http://www.w3.org/2000/svg" p-id="5096">
								<title>表情</title>
								<path d="M506.88 993.28C238.6944 993.28 20.48 775.0656 20.48 506.88S238.6944 20.48 506.88 20.48 993.28 238.6944 993.28 506.88 775.0656 993.28 506.88 993.28z m0-900.7104C278.4256 92.5696 92.5696 278.4256 92.5696 506.88S278.4256 921.1904 506.88 921.1904 921.1904 735.3344 921.1904 506.88 735.3344 92.5696 506.88 92.5696z" p-id="5097" />
								<path
									d="M522.5472 800.256c-105.8816 0-200.704-58.88-241.5616-150.016l65.7408-29.4912c29.2864 65.3312 98.304 107.52 175.8208 107.52 99.7376 0 146.2272-81.7152 161.5872-116.736l66.048 28.8768c-43.8272 100.0448-128.9216 159.8464-227.6352 159.8464zM335.7696 349.9008c29.7984 0 54.0672 24.1664 54.0672 54.0672v52.736c0 29.7984-24.1664 54.0672-54.0672 54.0672s-54.0672-24.1664-54.0672-54.0672v-52.736c0-29.9008 24.1664-54.0672 54.0672-54.0672zM689.5616 349.9008c29.7984 0 54.0672 24.1664 54.0672 54.0672v52.736c0 29.7984-24.1664 54.0672-54.0672 54.0672-29.7984 0-54.0672-24.1664-54.0672-54.0672v-52.736c0.1024-29.9008 24.2688-54.0672 54.0672-54.0672z"
									p-id="5098"
								/>
							</svg>
						</div>
					</div>
					<div class="wath-message" v-if="Message.Status">{{Message.Content}}</div>
				</form>
			</div>
		</div>
	</div>
</template>
<script>
export default {
	props: ["api", "title", "name", "avatar", "id", "logo"],
	data() {
		return {
			ws: null,
			enable: false,
			emojiEnable: false,
			menusEnable: false,
			content: "",
			ChatRecord: [],
			Id: this.id,
			Author: {
				Name: "",
				Avatar: "",
			},
			Message: {
				Status: false,
				Content: "",
			},
		}
	},
	created() {
		if (this.name == null || this.name == "") {
			if (localStorage.getItem("ChatAuthorName") == null) {
				localStorage.setItem("ChatAuthorName", "游客" + 'xxxxxxxxxxxx4xxxyxxxxxxxxxxxxxxx'.replace(/[xy]/g, function (c) {
					let r = Math.random() * 16 | 0, v = c == 'x' ? r : (r & 0x3 | 0x8);
					return v.toString(16);
				}));
			}
			this.Author.Name = localStorage.getItem("ChatAuthorName");
		} else {
			this.Author.Name = this.name;
		}
		if (this.avatar == null) {
			this.Author.Avatar = this.avatar;
		}
		if (this.id == null || this.id == "") {
			return;
		}
		this.Messages("正在连接服务器...");
		this.ws = new WebSocket(this.api + "?id=" + this.id + "&name=" + this.Author.Name + "&avatar=" + this.Author.Avatar);
		this.ws.onopen = () => {    //建立链接时触发
			this.WebSocketSend("ping")
		}
		this.ws.onmessage = (res) => {  // 接收到消息时触发
			this.WebSocketGet(res)
		}
		this.ws.onclose = (res) => {
			this.WebSocketSend("已断开连接...")
		}
		this.ws.onerror = (res) => {
			this.WebSocketSend("已与服务器断开连接...")
		}
	},
	methods: {
		WebSocketSend(content) {    //发送消息
			if (this.ws.readyState != 1) {
				return
			}
			this.ws.send(content);
		},
		WebSocketGet(res) { //接收消息处理
			if (this.ws.readyState != 1) {
				return
			}
			let handleMessage = JSON.parse(res.data)
			if (handleMessage.message == "pong") {
				this.Messages("已连接服务器...", true);
				this.ChatRecord = handleMessage.data ? handleMessage.data : [];
			} else {
				this.ChatRecord.push(handleMessage.data)
				this.$refs.showitems.scrollTop = 10000;
			}
		},
		UpdateContent() {
			if (this.content.length == 0) {
				this.Messages("消息内容不能为空...", true)
				return;
			}
			this.WebSocketSend(this.content);
			this.content = "";
		},
		UpdateCommentsContent(e) {
			var start = this.$refs.editarea.selectionStart;
			this.content = this.content.slice(0, start) + e + this.content.slice(start, this.$refs.editarea.length)
			this.emojiEnable = false;
			this.$refs.editarea.focus();
			return;
		},
		Messages(e, status) {
			this.Message.Status = true;
			this.Message.Content = e;
			if (status) {
				setTimeout(() => {
					this.Message.Status = false;
					this.Message.Content = ""
				}, 3000);
			}
		}
	},
}
</script>
<style scoped>
li {
	list-style: none;
}
.placeholder {
	-ms-flex-positive: 1;
	-webkit-box-flex: 1;
	flex-grow: 1;
}
a {
	text-decoration: none;
}

.chat-enable {
	position: fixed;
	z-index: 99;
	right: 10px;
	width: 75px;
	bottom: 10px;
}

img.chat-logo {
	max-width: 100%;
	width: 100%;
	cursor: pointer;
}
.chat-wath {
	position: fixed;
	z-index: 9999;
	display: flex;
	flex-wrap: wrap;
	max-width: 500px;
	width: 100%;
	height: 600px;
	margin: auto;
	left: 0;
	right: 0;
	top: 0;
	bottom: 0;
	background: #fff;
	box-shadow: 0 2px 12px 0 rgba(0, 0, 0, 0.14), 0 3px 1px -2px rgba(0, 0, 0, 0.2), 0 1px 5px 0 rgba(0, 0, 0, 0.12);
	border-radius: 5px;
}
.chat-wath * {
	-webkit-box-sizing: border-box;
	-webkit-tap-highlight-color: transparent;
	box-sizing: border-box;
	word-wrap: break-word;
}
.wath-header {
	-webkit-box-flex: 0;
	flex: 0 0 100%;
	height: 50px;
	line-height: 50px;
	border-radius: 5px 5px 0 0;
	display: flex;
	flex-wrap: wrap;
	background: #00a1d6;
}
.wath-header-f,
.wath-header-r {
	-webkit-box-flex: 0;
	flex: 0 0 100%;
	max-width: 50px;
	text-align: center;
	height: 100%;
}

.wath-header-t {
	-webkit-box-flex: 0;
	flex: 0 0 100%;
	max-width: calc(100% - 100px);
	text-align: center;
	height: 100%;
}

.wath-header-t span {
	font-size: 14px;
	font-weight: 600;
	color: #ffffff;
	padding: 0 10px;
}

.wath-header svg {
	width: 50px;
	height: 50px;
	padding: 15px 5px;
	cursor: pointer;
}
.wath-header svg path {
	fill: #fff;
}
.wath-header svg:hover {
	background: rgba(255, 255, 255, 0.0784313725490196);
}
.wath-header-menus {
	position: absolute;
	top: 50px;
	bottom: 50px;
	width: 50%;
	background-color: #363e47;
	left: 0;
	transition: 0.3s;
	z-index: 1;
}
li.wath-header-menus-list {
	height: 50px;
	line-height: 50px;
	color: #fff;
	padding: 0 10px;
	text-align: center;
	width: 100%;
	font-size: 1rem;
	text-overflow: ellipsis;
	overflow: hidden;
	white-space: nowrap;
	vertical-align: middle;
	border-bottom: 1px solid rgba(255, 255, 255, 0.04);
	cursor: pointer;
}
li.wath-header-menus-list:hover {
	background-color: #3f4953;
}
a.wath-theme {
	position: absolute;
	bottom: 0;
	width: 100%;
	margin: 0;
	padding: 0;
	color: #fff;
	font-size: 10px;
	text-align: center;
	display: block;
}

.wath-content {
	-webkit-box-flex: 0;
	flex: 0 0 100%;
	max-width: 100%;
	height: calc(100% - 100px);
	background: #f4f4f4;
	padding: 10px;
	overflow: auto;
}
.wath-content::-webkit-scrollbar {
	width: 5px;
	height: 5px;
}
.wath-content::-webkit-scrollbar-button {
	width: 0;
	height: 0;
}
.wath-content::-webkit-scrollbar-thumb {
	background: #929292;
}

li.wath-record-list {
	width: 100%;
	display: flex;
	flex-wrap: wrap;
	padding: 10px 0;
	position: relative;
}
.wath-record-list-author {
	display: flex;
	flex-wrap: wrap;
	-webkit-box-flex: 0;
	flex: 0 0 100%;
	max-width: 100%;
	height: 30px;
	line-height: 30px;
}
.author-avatar img.avatar {
	max-width: 30px;
	max-height: 30px;
	width: 100%;
	height: 100%;
	border-radius: 100%;
	border: 1px solid #ddd;
}
.author-nickname {
	max-width: 70%;
	margin: 0 10px;
	text-overflow: ellipsis;
	overflow: hidden;
	white-space: nowrap;
	vertical-align: middle;
	font-weight: 600;
	font-size: 13px;
	color: #757575;
	cursor: pointer;
}
.author-nickname:hover {
	color: #ff6666;
}
.author-date {
	font-size: 10px;
	color: #777;
	padding: 0 10px;
}

.wath-record-list-contens {
	-webkit-box-flex: 0;
	flex: 0 0 100%;
	max-width: 100%;
	min-height: 30px;
	padding: 5px 50px;
}
.wath-record-list-contens p {
	background: #fff;
	max-width: 100%;
	min-width: 50px;
	min-height: 30px;
	padding: 10px;
	font-size: 14px;
	color: #626c76;
	margin: 0;
}
.wath-record-list-contens.itself p {
	float: right;
	border-radius: 10px 0px 10px 10px;
}
.wath-record-list-contens.they p {
	float: left;
	border-radius: 0 10px 10px 10px;
}
.wath-record-list-contens.itself:after {
	content: "";
	width: 0;
	height: 0;
	border-style: solid;
	border-width: 15px 20px 0 0;
	border-color: #fff transparent transparent transparent;
	position: absolute;
	right: 30px;
}
.wath-record-list-contens.they:after {
	content: "";
	width: 0;
	height: 0;
	border-style: solid;
	border-width: 0 20px 15px 0;
	border-color: transparent #fff transparent transparent;
	position: absolute;
	left: 30px;
}
.wath-footer {
	position: relative;
	-webkit-box-flex: 0;
	flex: 0 0 100%;
	height: 50px;
	line-height: 50px;
	border-top: 1px solid #e3e3e3;
	background-color: #fff;
	border-radius: 0 0 5px 5px;
}
form.wath-form {
	margin: 0;
	display: flex;
	flex-wrap: wrap;
	height: 100%;
	width: 100%;
}
.wath-edit {
	width: 100%;
	height: 100%;
	display: flex;
	flex-wrap: wrap;
	padding: 5px;
	position: relative;
}
.edit-textarea {
	-ms-flex-positive: 1;
	-webkit-box-flex: 1;
	flex-grow: 1;
	height: 100%;
}
.edit-textarea textarea {
	display: block;
	font-size: 10px;
	width: 100%;
	height: 100%;
	border: none;
	text-align: left;
	transition: all 0.3s linear;
	color: #555;
	background-color: #ececec;
	border-radius: 7px;
	padding: 10px;
	outline: none;
	font-family: -webkit-body;
	cursor: pointer;
}
.edit-push,
.edit-emoji {
	-webkit-box-flex: 0;
	flex: 0 0 100%;
	max-width: 10%;
	height: 100%;
	text-align: center;
}
button.edit-push {
	border: none;
	background: none;
}
button.edit-push:hover {
	outline: none;
}
.wath-edit svg {
	width: 100%;
	height: 100%;
	cursor: pointer;
	padding: 5px;
}

.wath-edit svg path {
	fill: rgba(70, 68, 68, 0.7098039215686275);
}
.wath-edit svg:hover path {
	fill: #00a1d6;
}
.emoji-item {
	position: absolute;
	width: 100%;
	right: 0;
	background: #fff;
	padding: 5px;
	bottom: 0px;
	opacity: 0;
	z-index: -1;
	transition: all 0.3s;
}
.emoji-item.on {
	opacity: 1;
	z-index: inherit;
	transition: all 0.3s;
	bottom: 50px;
}
ul.emoji-box {
	list-style: none;
	margin: 0;
	padding: 0;
	position: relative;
	display: flex;
	flex-wrap: wrap;
}

a.emoji-list {
	-webkit-box-flex: 0;
	flex: 0 0 100%;
	max-width: 7%;
	font-size: 16px;
	padding: 0 4px;
	line-height: 40px;
	border: 1px solid #d9d9d9;
	cursor: pointer;
}

a.emoji-list img {
	max-width: 100%;
	max-height: 100%;
	width: 100%;
	vertical-align: middle;
	object-fit: cover;
}

.wath-message {
	position: absolute;
	top: -30px;
	height: 30px;
	line-height: 30px;
	background: #fff;
	width: 100%;
	text-align: center;
	font-size: 14px;
	color: #757575;
}
</style>

<style scoped>
@media (max-width: 600px) {
	.chat-wath {
		max-width: 90%;
		max-height: 500px;
	}
	.author-nickname {
		max-width: 50%;
	}
	a.emoji-list {
		max-width: 13.333%;
	}
}
</style>