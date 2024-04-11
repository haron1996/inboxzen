<script lang="ts">
	import { onMount } from 'svelte';
	import { URL } from '../../store';
	import { updateErrorMessages, updateSuccessMessages } from '../../utils';
	import Button from '../Button.svelte';

	interface Keyword {
		keyword?: string | undefined;
		date_added?: string | undefined;
		email_address?: string | undefined;
	}

	let keyword: Keyword = {};
	let keywords: Keyword[] = [];

	function handleKeywordInputBlur() {
		if (keyword.keyword === undefined) return;

		let found = false;

		keywords.forEach((kw) => {
			if (kw.keyword === keyword.keyword) {
				found = true;
				return;
			}
		});

		if (found) {
			keyword = {};
			return;
		}

		keywords = [keyword, ...keywords];

		keyword = {};
	}

	function handleKeywordInputKeydown(e: KeyboardEvent) {
		if (e.code !== 'Enter' && e.code !== 'Comma') return;

		e.preventDefault();

		handleKeywordInputBlur();
	}

	async function updateKeywords() {
		const url = `${$URL}/private/updatekeywords`;

		const response = await fetch(url, {
			method: 'POST',
			mode: 'cors',
			cache: 'no-cache',
			credentials: 'include',
			headers: {
				'Content-Type': 'application/json'
			},
			redirect: 'follow',
			referrerPolicy: 'no-referrer',
			body: JSON.stringify({
				keywords: keywords.map((kw) => kw.keyword)
			})
		});

		const result = await response.json();

		if (!response.ok) {
			const message = result.message;
			switch (message) {
				case 'The access token provided is invalid or has expired. Please log in again.':
					updateErrorMessages(message);
					setTimeout(() => {
						location.href = '/';
					}, 3000);
					return;
				default:
					updateErrorMessages(message);
					return;
			}
		}

		keywords = result;

		updateSuccessMessages('Keywords updated successfully');
	}

	async function getKeywords() {
		const url = `${$URL}/private/getkeywords`;

		const response = await fetch(url, {
			method: 'GET',
			mode: 'cors',
			cache: 'no-cache',
			credentials: 'include',
			headers: {
				'Content-Type': 'application/json'
			},
			redirect: 'follow',
			referrerPolicy: 'no-referrer'
		});

		const result = await response.json();

		if (!response.ok) {
			const message = result.message;
			switch (message) {
				case 'The access token provided is invalid or has expired. Please log in again.':
					updateErrorMessages(message);
					setTimeout(() => {
						location.href = '/';
					}, 3000);
					return;
				default:
					updateErrorMessages(message);
					return;
			}
		}

		keywords = result;
	}

	onMount(() => {
		getKeywords();
	});
</script>

<div class="vip-keywords">
	<div class="top">
		<p>VIP KEYWORDS</p>
		<span>Emails containing any of these keywords will be delivered instantly</span>
	</div>
	<div class="keywords">
		{#each keywords as { keyword }}
			<div class="keyword">
				<span>{keyword}</span>
				<svg
					width="24px"
					height="24px"
					stroke-width="1.5"
					viewBox="0 0 24 24"
					fill="none"
					xmlns="http://www.w3.org/2000/svg"
					color="#000000"
					><path
						d="M6.75827 17.2426L12.0009 12M17.2435 6.75736L12.0009 12M12.0009 12L6.75827 6.75736M12.0009 12L17.2435 17.2426"
						stroke="#000000"
						stroke-width="1.5"
						stroke-linecap="round"
						stroke-linejoin="round"
					></path></svg
				>
			</div>
		{/each}
		<input
			type="text"
			name="keyword"
			id="keyword"
			autocomplete="off"
			placeholder="add keyword..."
			bind:value={keyword.keyword}
			on:blur={handleKeywordInputBlur}
			on:keydown={handleKeywordInputKeydown}
		/>
	</div>
	<div class="bottom">
		<Button
			height={4}
			width={10}
			backgroundColor="#00a6fb"
			borderRadius={0.6}
			color="rgb(255, 255, 255)"
			padding={0.5}
			text="update"
			onClick={updateKeywords}
		/>
	</div>
</div>

<style lang="scss">
	.vip-keywords {
		width: 80rem;
		min-height: max-content;
		background-color: $white;
		display: flex;
		flex-direction: column;
		border-radius: 0.3rem;
		padding: 2rem;
		gap: 2rem;
		border: 0.1rem solid $black-4;
		box-shadow: rgba(0, 0, 0, 0.16) 0rem 0.1rem 0.4rem;

		.top {
			display: flex;
			flex-direction: column;
			justify-content: center;
			gap: 0.5rem;

			p {
				font-family: $spline;
				font-size: 1.1rem;
				text-transform: uppercase;
				font-weight: 500;
			}

			span {
				font-size: 1.3rem;
				font-family: $spline;
				font-weight: 500;
				color: $black-1;
			}
		}

		.keywords {
			min-height: max-content;
			display: flex;
			flex-flow: row wrap;
			align-content: start;
			align-items: center;
			gap: 1rem;

			.keyword {
				background-color: #f6f5f5;
				min-height: max-content;
				display: flex;
				align-items: center;
				padding: 0.5rem 0.5rem 0.5rem 0.7rem;
				width: max-content;
				border-radius: 0.5rem;
				gap: 0.3rem;
				box-shadow:
					rgba(0, 0, 0, 0.05) 0rem 0rem 0rem 0.1rem,
					rgb(209, 213, 219) 0rem 0rem 0rem 0.1rem inset;

				span {
					font-size: 1.3rem;
					font-family: $spline;
					font-weight: 500;
					color: $black-1;
				}

				svg {
					height: 1.5rem;
					width: 1.5rem;
					cursor: pointer;
				}
			}

			input {
				border: none;
				outline: none;
			}
		}

		.bottom {
			display: flex;
			align-items: center;
			justify-content: flex-end;
		}
	}
</style>
