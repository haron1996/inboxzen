interface UserEmailAccount {
	id?: string;
	account_name?: string;
	date_added?: string;
	email_address?: string;
	primaryaccount?: boolean;
	profile_picture?: string;
	user_id?: string;
	running?: boolean;
}

// user session interface
export interface Session {
	email?: UserEmailAccount | undefined;
	emails?: UserEmailAccount[] | undefined;
}

// vip domains interface
export interface Domain {
	id?: string | undefined;
	date_added?: string | undefined;
	domain_name?: string | undefined;
	email_address?: string | undefined;
}

// vip email addresses interface
export interface Email {
	id?: string | undefined;
	vip_email_address?: string | undefined;
	date_added?: string | undefined;
	email_address?: string | undefined;
}

// vip keywords interface
export interface Keyword {
	id?: string | undefined;
	keyword?: string | undefined;
	date_added?: string | undefined;
	email_address?: string | undefined;
}

// inbox delivey times interface
export interface Time {
	id?: string | undefined;
	delivery_time?: string | undefined;
}
