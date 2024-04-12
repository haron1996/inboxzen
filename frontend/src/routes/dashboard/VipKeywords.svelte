<script lang="ts">
	import { keyword, keywords, loading } from '../../store';
	import { updateVipKeywords } from '../../utils';
	import Button from '../Button.svelte';
	import CardSkeleton from './CardSkeleton.svelte';

	function handleKeywordInputBlur() {
		if ($keyword.keyword === undefined) return;

		let found = false;

		$keywords.forEach((kw) => {
			if (kw.keyword === $keyword.keyword) {
				found = true;
				return;
			}
		});

		if (found) {
			$keyword = {};
			return;
		}

		$keywords = [$keyword, ...$keywords];

		$keyword = {};
	}

	function handleKeywordInputKeydown(e: KeyboardEvent) {
		if (e.code !== 'Enter' && e.code !== 'Comma') return;

		e.preventDefault();

		handleKeywordInputBlur();
	}

	function removeVIPKeyword(e: MouseEvent) {
		const t = e.currentTarget as HTMLElement;

		const div = t.closest('.keyword') as HTMLDivElement | null;

		if (div === null) return;

		const keywrd = div.innerText;

		$keywords = $keywords.filter((kw) => kw.keyword !== keywrd);
	}
</script>

<div class="vip-keywords">
	<div class="top">
		<p>VIP KEYWORDS</p>
		<span>Emails containing any of these keywords will be delivered instantly</span>
	</div>
	<div class="keywords">
		{#if $loading}
			<CardSkeleton height={10} width={80} padding={0} borderRadius={0} />
		{/if}
		{#each $keywords as { keyword }}
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
					role="none"
					on:click={removeVIPKeyword}
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
			bind:value={$keyword.keyword}
			on:blur={handleKeywordInputBlur}
			on:keydown={handleKeywordInputKeydown}
		/>
	</div>
	<div class="bottom">
		<Button
			height={4}
			width={10}
			borderRadius={0.3}
			color="rgb(255, 255, 255)"
			padding={0.5}
			text="update"
			onClick={() => {
				updateVipKeywords($keywords);
			}}
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
				font-family: $spline;
				padding: 0.5rem;
			}
		}

		.bottom {
			display: flex;
			align-items: center;
			justify-content: flex-end;
			padding: 1rem;
		}
	}
</style>
