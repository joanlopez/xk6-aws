import {check} from 'k6';
import {
	createEventBus,
	deleteEventBus,
	listEventBuses,
	putEvents,
	putRule,
	putTargets,
	removeTargets,
} from 'k6/x/aws';

// Set your k6 run configuration:
// https://k6.io/docs/using-k6/k6-options
export const options = {
	iterations: 1,

	// Demonstrative k6 thresholds.
	thresholds: {
		checks: [{threshold: 'rate == 1.00', abortOnFail: true}],
	},
};

export default function () {
	// Create an event bus.
	const eventBusName = 'test-event-bus';
	const {event_bus_arn} = createEventBus({name: eventBusName});
	check(event_bus_arn, {
		'event bus creation must succeed': (event_bus_arn) => event_bus_arn.includes(`event-bus/${eventBusName}`),
	});

	// List event buses to confirm creation. It must return two event buses:
	// - the default event bus
	// - the one we created
	const {event_buses} = listEventBuses();
	check(event_buses, {
		'it must return two event buses': (buses) => buses.length === 2,
		'test-event-bus must be in the list': (buses) => buses.some(b => b.name.normalize() === eventBusName),
	});

	// Put an event onto the event bus.
	const eventDetail = JSON.stringify({key1: 'value1'});
	const putResponse = putEvents({
		entries: [
			{
				event_bus_name: eventBusName,
				source: 'my.source',
				detail_type: 'myDetailType',
				detail: eventDetail,
			},
		],
	});

	// Check that the event was successfully put.
	check(putResponse, {
		'put events must succeed': (res) => res.failed_entry_count === 0,
	});

	// Create a rule.
	const ruleName = 'test-rule';
	const {rule_arn} = putRule({
		name: ruleName,
		event_bus_name: eventBusName,
		event_pattern: JSON.stringify({
			source: ['my.source'],
		}),
	});

	check(rule_arn, {
		'rule creation must succeed': (arn) => arn.includes(`rule/${eventBusName}/${ruleName}`),
	});

	// Put a target for the rule.
	const targetId = 'test-target';
	const targetArn = 'arn:aws:lambda:us-east-1:123456789012:function:test-function';
	const putTargetsResponse = putTargets({
		rule: ruleName,
		event_bus_name: eventBusName,
		targets: [
			{
				id: targetId,
				arn: targetArn,
			},
		],
	});

	check(putTargetsResponse, {
		'put targets must succeed': (res) => res.failed_entry_count === 0,
	});

	// Remove the target.
	const removeTargetsResponse = removeTargets({
		rule: ruleName,
		event_bus_name: eventBusName,
		ids: [targetId],
	});

	check(removeTargetsResponse, {
		'remove targets must succeed': (res) => res.failed_entry_count === 0,
	});

	// Delete the event bus (idempotent).
	deleteEventBus({name: eventBusName});
}