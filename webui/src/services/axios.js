import axios from "axios";
import {readToken} from '../services/session'

const api = axios.create({
	baseURL: __API_URL__,
	timeout: 1000 * 5
});

api.interceptors.request.use(config => {
	config.headers['token'] = readToken();
	return config;
});

export default api;
