interface UserEmailAccount {
	account_name: string;
	date_added: string;
	email_address: string;
	primaryaccount: boolean;
	profile_picture: string;
	user_id: string;
}

// user session interface
export interface Session {
	email?: UserEmailAccount | undefined;
	emails?: UserEmailAccount[] | undefined;
}

// vip domains interface
export interface Domain {
	date_added?: string | undefined;
	domain_name?: string | undefined;
	email_address?: string | undefined;
}

// vip email addresses interface
export interface Email {
	vip_email_address?: string | undefined;
	date_added?: string | undefined;
	email_address?: string | undefined;
}

// vip keywords interface
export interface Keyword {
	keyword?: string | undefined;
	date_added?: string | undefined;
	email_address?: string | undefined;
}

// inbox delivey times interface
export interface Time {
	hour?: string | undefined;
	minutes?: string | undefined;
	am_pm?: string | undefined;
}
