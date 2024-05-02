class Auth {
	constructor() {
		this.sessionStorageUser = 'scortiousprime';
		this.sessionStorageRoom = 'rapidchatRoom';
	}

	login(name, roomName, cb) {
		var userId = this.createGuid();
		sessionStorage.setItem(this.sessionStorageUser, JSON.stringify({
			_name: name,
			_userId: userId
		}));
		sessionStorage.setItem(this.sessionStorageRoom, JSON.stringify({
			_roomName: roomName
		}));
		cb();
	}

	logout(cb) {
		sessionStorage.removeItem(this.sessionStorageUser);
		sessionStorage.removeItem(this.sessionStorageRoom);
		cb();
	}

	isAuthenticated() {
		var user = sessionStorage.getItem(this.sessionStorageUser);
		var room = sessionStorage.getItem(this.sessionStorageRoom);
		return user && room; 
	}

	getUserName() {
		return JSON.parse(sessionStorage.getItem(this.sessionStorageUser))._name;
	}

	getUserId() {
		return JSON.parse(sessionStorage.getItem(this.sessionStorageUser))._userId;
	}

	getRoomName() {
		return JSON.parse(sessionStorage.getItem(this.sessionStorageRoom))._roomName;
	}

	createGuid() {
		function S4() {
			return (((1 + Math.random()) * 0x10000) | 0).toString(16).substring(1);
		}
		return (S4() + S4() + "-" + S4() + "-4" + S4().substring(0, 3) + "-" + S4() + "-" + S4() + S4() + S4()).toLowerCase();
	}
}

export default new Auth();
