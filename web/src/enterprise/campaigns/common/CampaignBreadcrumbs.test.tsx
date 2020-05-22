import React from 'react'
import { shallow } from 'enzyme'
import { CampaignBreadcrumbs } from './CampaignBreadcrumbs'

describe('CampaignBreadcrumbs', () => {
    test('new', () => expect(shallow(<CampaignBreadcrumbs />)).toMatchSnapshot())

    test('existing', () =>
        expect(
            shallow(<CampaignBreadcrumbs campaign={{ name: 'My campaign', url: 'https://example.com' }} />)
        ).toMatchSnapshot())
})
