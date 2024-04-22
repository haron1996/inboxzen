<script lang="ts">
	import { afterNavigate } from '$app/navigation';
	import { getFirstTimeSenders, getUserAccount, moveEmailsToInbox } from '../../../utils';

	afterNavigate(() => {
		getUserAccount();
		//getFirstTimeSenders();
		moveEmailsToInbox();
	});
</script>

<svelte:head>
	<title>The Screener | Gmail Inbox Zero</title>
</svelte:head>

<section>
	<div class="new-senders" id="new-senders">
		<!-- <div class="top">
			<p>THE PEOPLE BELOW ARE EMAILING YOU FOR THE FIRST TIME. WANT TO HEAR FROM THEM?</p>
		</div> -->
	</div>
</section>

<style lang="scss">
	section {
		width: 100%;
		height: 100%;

		.new-senders {
			width: 80%;
			height: 95%;
			border-radius: 0.5rem;
			box-shadow: rgba(0, 0, 0, 0.09) 0px 3px 12px;
			margin: 1rem auto 0 auto;
			background-color: #ffffff;
			padding: 2rem;

			// .top {
			// 	p {
			// 		font-family: $spline;
			// 		text-transform: uppercase;
			// 		font-size: 1.1rem;
			// 		font-weight: 600;
			// 		color: #403d39;
			// 	}
			// }
		}
	}
</style>
