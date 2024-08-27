import axios from "axios";
import {readToken} from '../services/session'
import router from "../router/index.js"

const api = axios.create({
	baseURL: __API_URL__,
	timeout: 1000 * 5
});

api.interceptors.request.use(config => {
	config.headers['token'] = readToken();
	return config;
});

api.interceptors.response.use(function (response) {
	return response;
}, function (error) {
	if (error.response) {
		if (error.response.status == 403) {
			router.push("/error?code=403");
		} else if (error.response.status == 401) {
			writeUser();
			router.push("/login");
		} else {
			return Promise.reject(error);
		}
	} else {
		if (error.request) {
			console.log("No Response:\n\t", error.request);
		} else {
			console.log("Error:\n\t", error);
		}
		return Promise.reject(error);
	}
});

export default api;
