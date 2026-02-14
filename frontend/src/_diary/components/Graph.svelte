<script lang="ts">
	import Tooltip, { Wrapper, Title, Content, Link, RichActions } from '@smui/tooltip'
	import { getContext } from 'svelte'
	import type { Writable } from 'svelte/store'
	import GraphTooltip from './GraphTooltip.svelte'

	const local: Writable<any> = getContext('local')
	const remote: Writable<any> = getContext('remote')
	const action: Writable<any> = getContext('action')
</script>

<div class="graph">
	<div class="wrapper">
		<div class="heading">
			<h3>Activity Graph</h3>
			<Wrapper rich style="flex: 1">
				<span class="material-symbols-outlined icon">help</span>
				<Tooltip>
					<Content style="width: 260px">
						3 aspects of daily score represented<br />in the color mixture of each squares.
						<div class="r-1">
							<div class="preview-square" style="background-color: dodgerblue" />
							Gain
							<span class="preview-desc">Experienced something new</span>
						</div>
						<div class="r-1">
							<div class="preview-square" style="background-color: forestgreen" />
							Emotional
							<span class="preview-desc">Feeling / Enjoyment / Mood</span>
						</div>
						<div class="r-1">
							<div class="preview-square" style="background-color: indianred" />
							Productivity
							<span class="preview-desc">Things accomplished</span>
						</div>
					</Content>
				</Tooltip>
			</Wrapper>
			<select on:change={(e) => $action.setYear(e.target.value)}>
				{#each $local.yearOptions as year}
					<option value={year} selected={year === $local.year}>{year}</option>
				{/each}
			</select>
		</div>
		<div class="graph">
			<div class="months">
				<span>Jan</span>
				<span>Feb</span>
				<span>Mar</span>
				<span>Apr</span>
				<span>May</span>
				<span>Jun</span>
				<span>Jul</span>
				<span>Aug</span>
				<span>Sep</span>
				<span>Oct</span>
				<span>Nov</span>
				<span>Dec</span>
			</div>
			<ul class="squares">
				{#each $remote.days as day, i}
					{#if i === 0}
						<div
							contenteditable
							class={`square ${$local.activeGraph?.day.date === day.date ? 'active' : ''}`}
							style={`background-color: #${day.graph?.toString(16).padStart(6, '0')};
								grid-row-start: ${$local.firstDayOffset}`}
							on:blur={() => $action.clearActiveGraph()}
							on:click={(ev) => $action.setActiveGraph(ev, day)}
							on:keypress={(ev) => $action.setActiveGraph(ev, day)}
						/>
					{:else}
						<div
							contentEditable
							class={`square ${$local.activeGraph?.day.date === day.date ? 'active' : ''}`}
							style={`background-color: #${day.graph?.toString(16).padStart(6, '0')}`}
							on:blur={() => $action.clearActiveGraph()}
							on:click={(ev) => $action.setActiveGraph(ev, day)}
							on:keypress={(ev) => $action.setActiveGraph(ev, day)}
						/>
					{/if}
				{/each}
			</ul>
		</div>
	</div>
	<GraphTooltip />
</div>

<style lang="scss">
	@import '../../styles/index';

	$gap: 6px;
	$square-size: 16px;
	$week-width: 22px;

	.wrapper {
		padding: 20px;
		margin: 20px;
		border: 1px #e1e4e8 solid;
		border-radius: 6px;
		width: calc(100% - 40px);
		max-width: 1156px;
		overflow-x: scroll;
		overflow-y: visible;
	}

	.heading {
		display: flex;
		justify-content: space-between;
		align-items: center;
		margin-bottom: 16px;
	}

	.icon {
		margin-left: 6px;
		font-size: 18px;
		@include tweak-clickable-icon($color-grey-800);
	}

	.preview-square {
		width: 14px;
		height: 14px;
		margin-right: 6px;
	}

	.preview-desc {
		margin: 0 0 0 6px;
		opacity: 0.6;
		font-size: 0.8em;
		line-height: 1em;
	}

	.graph {
		display: inline-grid;
		grid-template-areas:
			'months'
			'squares';
		grid-template-columns: auto 1fr;
		grid-gap: $gap;
	}

	.months {
		grid-area: months;

		display: grid;
		grid-template-columns:
			calc($week-width * 4) /* Jan */
			calc($week-width * 4) /* Feb */
			calc($week-width * 5) /* Mar */
			calc($week-width * 4) /* Apr */
			calc($week-width * 4) /* May */
			calc($week-width * 5) /* Jun */
			calc($week-width * 4) /* Jul */
			calc($week-width * 4) /* Aug */
			calc($week-width * 5) /* Sep */
			calc($week-width * 4) /* Oct */
			calc($week-width * 4) /* Nov */
			calc($week-width * 5) /* Dec */;

		> span {
			font-size: 14px;
		}
	}

	.squares {
		grid-area: squares;

		display: grid;
		grid-gap: $gap;
		grid-template-rows: repeat(7, $square-size);
		grid-auto-flow: column;
		grid-auto-columns: $square-size;
	}

	.square {
		border-radius: 2px;
		width: 100%;
		height: 100%;
		cursor: pointer;
		caret-color: transparent;
		border: 1px #a1a4a800 solid;

		&:hover {
			border: 1px #a1a4a8aa solid;
		}

		&:active,
		&.active {
			border: 1px #a1a4a8ff solid;
		}
	}
</style>
