interface account {
	account_name: string;
	date_added: string;
	email: string;
	primaryaccount: boolean;
	profile_picture: string;
	user_id: string;
}

export interface session {
	account?: account;
	accounts?: account[];
	token?: string;
}
