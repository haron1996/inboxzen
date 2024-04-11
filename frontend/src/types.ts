interface UserEmailAccount {
	account_name: string;
	date_added: string;
	email_address: string;
	primaryaccount: boolean;
	profile_picture: string;
	user_id: string;
}

export interface Session {
	email?: UserEmailAccount;
	emails?: UserEmailAccount[];
}
